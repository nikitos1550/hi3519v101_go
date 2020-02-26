require.config({
  paths: {
    bootstrap: './vendor/bootstrap.min',
    diffMatchPatch: './vendor/diff_match_patch.min',
    handlebars: './vendor/handlebars.min',
    handlebarsExtended: './utils/handlebars_helper',
    jquery: './vendor/jquery.min',
    locales: './locales/locale',
    lodash: './vendor/lodash.min',
    pathToRegexp: './vendor/path-to-regexp/index',
    prettify: './vendor/prettify/prettify',
    utilsSampleRequest: './utils/send_sample_request',
    ripples: './vendor/bootstrap-material-design/dist/js/ripples',
    materialDesign: './vendor/bootstrap-material-design/dist/js/material',
    list: './vendor/list.min'
  },
  shim: {
    bootstrap: {
      deps: ['jquery']
    },
    diffMatchPatch: {
      exports: 'diff_match_patch'
    },
    handlebars: {
      exports: 'Handlebars'
    },
    handlebarsExtended: {
      deps: ['jquery', 'handlebars'],
      exports: 'Handlebars'
    },
    prettify: {
      exports: 'prettyPrint'
    },
    ripples: {
      deps: ['jquery']
    },
    materialDesign: {
      deps: ['jquery']
    },
  },
  urlArgs: 'v=' + (new Date()).getTime(),
  waitSeconds: 15
});

require([
  'jquery',
  'lodash',
  'locales',
  'handlebarsExtended',
  './api_project.js',
  './api_data.js',
  'prettify',
  'utilsSampleRequest',
  'bootstrap',
  'pathToRegexp',
  'ripples',
  'materialDesign',
  'list'
], function ($, _, locale, Handlebars, apiProject, apiData, prettyPrint, sampleRequest) {
  var api = {};
  var templateHeader;
  var templateFooter;
  var templateArticle;
  var templateGenerator;
  var templateProject;
  var templateSections;
  var templateSidenav;
  var apiGroups = {};
  var apiGroupTitles = {};
  var nav = [];
  // load google web fonts
  loadGoogleFontCss();
  renderPage(apiData.api);

  function renderPage(data) {
    initDocument();
    api = sortGroup(data);
    apiGroups = initGroups();
    addHeader();
    initNavigationList(apiGroups);
    AddFooter();
    documentTitle();
    sideNav();
    initContent();
    initDynamic();
    // remove loader
    $('#loader').remove();
    clickSideNav();
    menuClick();
  }

  function bindSearchResult(data) {
    apiGroups = {};
    nav = [];
    api = sortGroup(data);
    apiGroups = initGroups();
    addHeader();
    initNavigationList(apiGroups);
    AddFooter();
    sideNav();
    initContent();
    initDynamic();
    $('#loader').remove();
    clickSideNav();
    menuClick();
  }

  function initDocument() {
    // template
    templateHeader = Handlebars.compile($('#template-header').html());
    templateFooter = Handlebars.compile($('#template-footer').html());
    templateArticle = Handlebars.compile($('#template-article').html());
    templateGenerator = Handlebars.compile($('#template-generator').html());
    templateProject = Handlebars.compile($('#template-project').html());
    templateSections = Handlebars.compile($('#template-sections').html());
    templateSidenav = Handlebars.compile($('#template-sidenav').html());
    // apiProject defaults
    if (!apiProject.template) { apiProject.template = {}; }
    if (apiProject.template.forceLanguage) { locale.setLanguage(apiProject.template.forceLanguage); }
    // Setup jQuery Ajax
    $.ajaxSetup(apiProject.template.jQueryAjaxSetup);
  }

  function sortGroup(api) {
    // grouped by group
    var apiByGroup = _.groupBy(api, function (entry) {
      return entry.group;
    });
    // grouped by group and name
    var apiByGroupAndName = {};
    $.each(apiByGroup, function (index, entries) {
      apiByGroupAndName[index] = _.groupBy(entries, function (entry) {
        return entry.name;
      });
    });
    // sort api within a group by title ASC and custom order
    var newList = [];
    var umlauts = { 'ä': 'ae', 'ü': 'ue', 'ö': 'oe', 'ß': 'ss' };
    $.each(apiByGroupAndName, function (index, groupEntries) {
      // get titles from the first entry of group[].name[]
      var titles = [];
      $.each(groupEntries, function (titleName, entries) {
        var title = entries[0].title;
        if (title !== undefined) {
          title.toLowerCase().replace(/[äöüß]/g, function ($0) { return umlauts[$0]; });
          titles.push(title + '#~#' + titleName); // '#~#' keep reference to titleName after sorting
        }
      });
      // sort by name ASC
      titles.sort();
      // custom order
      if (apiProject.order) { titles = sortByOrder(titles, apiProject.order, '#~#'); }
      // add single elements to the new list
      titles.forEach(function (name) {
        var values = name.split('#~#');
        var key = values[1];
        groupEntries[key].forEach(function (entry) {
          newList.push(entry);
        });
      });
    });

    return newList;
  }

  function initGroups() {
    $.each(api, function (index, entry) {
      apiGroups[entry.group] = 1;
      apiGroupTitles[entry.group] = entry.groupTitle || entry.group;
    });
    // sort groups
    apiGroups = Object.keys(apiGroups);
    apiGroups.sort();
    // custom order
    if (apiProject.order) { apiGroups = sortByOrder(apiGroups, apiProject.order); }
    return apiGroups;
  }

  function initNavigationList(apiGroups) {
    apiGroups.forEach(function (group) {
      var subNav = [];
      // Submenu
      var oldName = '';
      api.forEach(function (entry) {
        if (entry.group === group) {
          switch (entry.type.toLowerCase()) {
            case 'post':
              label = 'info';
              break;
            case 'get':
              label = "success";
              break;
            case 'put':
              label = 'warning';
              break;
            case 'delete' || 'del':
              label = 'danger';
              break;
            default:
              label = 'default';
          }
          if (oldName !== entry.name) {
            subNav.push({
              title: entry.title,
              group: group,
              name: entry.name,
              type: entry.type,
              label: label
            });
          } else {
            subNav.push({
              title: entry.title,
              group: group,
              hidden: true,
              name: entry.name,
              type: entry.type,
              label: label
            });
          }
          oldName = entry.name;
        }
      });
      // Mainmenu entry
      nav.push({
        group: group,
        isHeader: true,
        title: apiGroupTitles[group],
        subNav: subNav
      });
      subNav = [];
    });
  }

  function add_nav(nav, content, index) {
    if (!content) return;
    var topics = content.match(/<h2>(.+?)<\/h2>/gi);
    topics.forEach(function (entry) {
      var title = entry.replace(/<.+?>/g, '');    // Remove all HTML tags for the title
      var entry_tags = entry.match(/id="(?:api-)?([^\-]+)-(.+)"/);    // Find the group and name in the id property
      var name = (entry_tags ? entry_tags[2] : null);
      if (title && name) {
        nav.splice(index, 0, {
          group: '_',
          name: name,
          //isHeader: false,
          title: title,
          //  isFixed: false,
        });
        index++;
      }
    });
  }

  function addHeader() {
    // Mainmenu Header entry
    if (apiProject.header) {
      nav.unshift({
        group: '_',
        isHeader: true,
        title: (apiProject.header.title == null) ? locale.__('General') : apiProject.header.title,
        //  isFixed: true
      });
      add_nav(nav, apiProject.header.content, 1);
    }
  }

  function AddFooter() {
    // Mainmenu Footer entry
    if (apiProject.footer && apiProject.footer.title != null) {
      nav.push({
        group: '_footer',
        isHeader: true,
        title: apiProject.footer.title,
        isFixed: true
      });
    }

  }

  function documentTitle() {
    // render page title
    var title = apiProject.title ? apiProject.title : 'apiDoc: ' + apiProject.name;
    $(document).attr('title', title);
  }

  function sideNav() {
    $('#sidenav').empty();
    $('#generator').empty();
    $('#project').empty();
    // render sidenav
    var fields = {
      nav: nav
    };
    $('#sidenav').append(templateSidenav(fields));

    // render Generator
    $('#generator').append(templateGenerator(apiProject));

    $('#project').append(templateProject(apiProject));

    // render apiDoc, header/footer documentation
    if (apiProject.header) {
      $('#header').empty();
      $('#header').append(templateHeader(apiProject.header));
    }

    if (apiProject.footer) {
      $('#footer').empty();
      $('#footer').append(templateFooter(apiProject.footer));
    }
  }

  function initContent() {
    // Render Sections and Articles
    var content = '';
    apiGroups.forEach(function (groupEntry) {
      var articles = [];
      var oldName = '';
      var fields = {};
      var title = groupEntry;
      var description = '';
      var label = "";

      // render all articles of a group
      api.forEach(function (entry) {
        if (groupEntry === entry.group) {
          if (oldName !== entry.name) {
            fields = {
              article: entry
            };
          } else {
            fields = {
              article: entry,
              hidden: true
            };
          }
          switch (entry.type.toLowerCase()) {
            case 'post':
              label = 'info';
              break;
            case 'get':
              label = "success";
              break;
            case 'put':
              label = 'warning';
              break;
            case 'delete' || 'del':
              label = 'danger';
              break;
            default:
              label = 'default';
          }
          fields.article.label = label;
          // add prefix URL for endpoint
          if (apiProject.url) {
            fields.article.url = fields.article.url.includes(apiProject.url) ? fields.article.url : apiProject.url + fields.article.url;
          }
          addArticleSettings(fields, entry);
          if (entry.groupTitle) { title = entry.groupTitle; }
          if (entry.groupDescription) { description = entry.groupDescription; }
          articles.push({
            article: templateArticle(fields),
            group: entry.group,
            name: entry.name
          });
          oldName = entry.name;
        }
      });
      // render Section with Articles
      var fields = {
        group: groupEntry,
        title: title,
        description: description,
        articles: articles
      };
      content += templateSections(fields);
    });

    $('#sections').empty();
    $('#sections').append(content);
  }

  /**
   * Check if Parameter (sub) List has a type Field.
   * Example: @apiSuccess          varname1 No type.
   *          @apiSuccess {String} varname2 With type.
   *
   * @param {Object} fields
   */
  function _hasTypeInFields(fields) {
    var result = false;
    $.each(fields, function (name) {
      if (_.any(fields[name], function (item) { return item.type; }))
        result = true;
    });
    return result;
  }

  /**
   * On Template changes, recall plugins.
   */
  function initDynamic() {
    // bootstrap popover
    $('a[data-toggle=popover]')
      .popover()
      .click(function (e) {
        e.preventDefault();
      });

    $('#sidenav li').removeClass('is-new');
    $('#scrollingNav').affix({
      offset: {
        top: 100,
        bottom: function () {
          return (this.bottom = $('.footer').outerHeight(true))
        }
      }
    });
    // tabs
    $('.nav-tabs-examples a').click(function (e) {
      e.preventDefault();
      $(this).tab('show');
    });
    // sample request switch
    $('.sample-request-switch').click(function (e) {
      var name = '.' + $(this).attr('name') + '-fields';
      $(name).addClass('hide');
      $(this).parent().next(name).removeClass('hide');
    });

    // init modules
    sampleRequest.initDynamic();
  }

  /**
   * Add article settings.
   */
  function addArticleSettings(fields, entry) {
    // add unique id
    fields.id = fields.article.group + '-' + fields.article.name;
    fields.id = fields.id.replace(/\./g, '_');

    if (entry.header && entry.header.fields)
      fields._hasTypeInHeaderFields = _hasTypeInFields(entry.header.fields);

    if (entry.parameter && entry.parameter.fields)
      fields._hasTypeInParameterFields = _hasTypeInFields(entry.parameter.fields);

    if (entry.error && entry.error.fields)
      fields._hasTypeInErrorFields = _hasTypeInFields(entry.error.fields);

    if (entry.success && entry.success.fields)
      fields._hasTypeInSuccessFields = _hasTypeInFields(entry.success.fields);

    if (entry.info && entry.info.fields)
      fields._hasTypeInInfoFields = _hasTypeInFields(entry.info.fields);

    // add template settings
    fields.template = apiProject.template;
  }

  /**
   * Load google fonts.
   */
  function loadGoogleFontCss() {
    var host = document.location.hostname.toLowerCase();
    var protocol = document.location.protocol.toLowerCase();
    var googleCss = '//fonts.googleapis.com/css?family=Source+Code+Pro|Source+Sans+Pro:400,600,700';
    if (host == 'localhost' || !host.length || protocol === 'file:')
      googleCss = 'http:' + googleCss;

    $('<link/>', {
      rel: 'stylesheet',
      type: 'text/css',
      href: googleCss
    }).appendTo('head');
  }

  /**
   * Return ordered entries by custom order and append not defined entries to the end.
   * @param  {String[]} elements
   * @param  {String[]} order
   * @param  {String}   splitBy
   * @return {String[]} Custom ordered list.
   */
  function sortByOrder(elements, order, splitBy) {
    var results = [];
    order.forEach(function (name) {
      if (splitBy)
        elements.forEach(function (element) {
          var parts = element.split(splitBy);
          var key = parts[1]; // reference keep for sorting
          if (key == name)
            results.push(element);
        });
      else
        elements.forEach(function (key) {
          if (key == name)
            results.push(name);
        });
    });
    // Append all other entries that ar not defined in order
    elements.forEach(function (element) {
      if (results.indexOf(element) === -1)
        results.push(element);
    });
    return results;
  }

  window.page = window.location.hash || "#api-_";

  $(document).ready(function () {
    $.material.init();
    var url = window.location.toString();
    var id = url.split('#')[1];
    if (id) {
      $(document).scrollTop($(location.hash)[0].offsetTop);
    }
  });


  /**
   * click menu
   */
  function menuClick() {
    $(".menu li").click(function () {
      if ($(this).is(".active")) {
        return;
      }
      $(".menu li").not($(this)).removeClass("active");
      $(this).addClass("active");
    });
  }

  /**
   * Content-Scroll on Navigation click.
   */
  function clickSideNav() {
    $('.sidenav').find('a').on('click', function (e) {
      e.preventDefault();
      var id = $(this).attr('href');
      if ($(id).length > 0) {
        $('html,body').animate({ scrollTop: parseInt($(id).offset().top) }, 1000);
      }
      window.location.hash = $(this).attr('href');
      if ($(this).parent().is(".nav-header")) {
        $('.nav-header').removeClass('active');
        $(this).parent().addClass('active');
      }
    });
  }


  /**
   * Set initial focus to search input
   */
  $('#scrollingNav .sidenav-search input.search').focus();

  /**
   * Detect ESC key to reset search
   */
  $(document).keyup(function (e) {
    // reset
    if (e.keyCode === 27) {
      $('span.search-reset').click();
    }
    // search
    if (e.keyCode === 13) {
      if ($("#searchText").val().trim() == "") {
        bindSearchResult(apiData.api);
      } else {
        var apis = apiData.api;
        var searchText = $("#searchText").val().trim();
        datas = $.grep(apis, function (value, index) {
          return value.title.includes(searchText);
        });
        bindSearchResult(datas);
      }
    }
  });

  /**
   * Search reset
   */
  $('span.search-reset').on('click', function () {
    $('#searchText').val("").focus();
    bindSearchResult(apiData.api);
  });

});
