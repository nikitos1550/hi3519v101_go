/*
 * Camera API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name        string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method      string
	// Pattern is the pattern of the URI.
	Pattern     string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/api/v1/",
		Index,
	},

	{
		"Apiinfo",
		http.MethodGet,
		"/api/v1/",
		Apiinfo,
	},

	{
		"ChannelGet",
		http.MethodGet,
		"/api/v1/channel",
		ChannelGet,
	},

	{
		"ChannelIdEnableDelete",
		http.MethodDelete,
		"/api/v1/channel/:id/enable",
		ChannelIdEnableDelete,
	},

	{
		"ChannelIdEnablePost",
		http.MethodPost,
		"/api/v1/channel/:id/enable",
		ChannelIdEnablePost,
	},

	{
		"ChannelIdGet",
		http.MethodGet,
		"/api/v1/channel/:id",
		ChannelIdGet,
	},

	{
		"ChannelIdPatch",
		http.MethodPatch,
		"/api/v1/channel/:id",
		ChannelIdPatch,
	},

	{
		"DebugGet",
		http.MethodGet,
		"/api/v1/debug",
		DebugGet,
	},

	{
		"DebugMppGet",
		http.MethodGet,
		"/api/v1/debug/mpp",
		DebugMppGet,
	},

	{
		"EncoderGet",
		http.MethodGet,
		"/api/v1/encoder",
		EncoderGet,
	},

	{
		"EncoderH264Get",
		http.MethodGet,
		"/api/v1/encoder/h264",
		EncoderH264Get,
	},

	{
		"EncoderH264IdDelete",
		http.MethodDelete,
		"/api/v1/encoder/h264/:id",
		EncoderH264IdDelete,
	},

	{
		"EncoderH264IdGet",
		http.MethodGet,
		"/api/v1/encoder/h264/:id",
		EncoderH264IdGet,
	},

	{
		"EncoderH264IdPatch",
		http.MethodPatch,
		"/api/v1/encoder/h264/:id",
		EncoderH264IdPatch,
	},

	{
		"EncoderH264Post",
		http.MethodPost,
		"/api/v1/encoder/h264",
		EncoderH264Post,
	},

	{
		"EncoderH265Get",
		http.MethodGet,
		"/api/v1/encoder/h265",
		EncoderH265Get,
	},

	{
		"EncoderH265IdDelete",
		http.MethodDelete,
		"/api/v1/encoder/h265/:id",
		EncoderH265IdDelete,
	},

	{
		"EncoderH265IdGet",
		http.MethodGet,
		"/api/v1/encoder/h265/:id",
		EncoderH265IdGet,
	},

	{
		"EncoderH265IdPatch",
		http.MethodPatch,
		"/api/v1/encoder/h265/:id",
		EncoderH265IdPatch,
	},

	{
		"EncoderH265Post",
		http.MethodPost,
		"/api/v1/encoder/h265",
		EncoderH265Post,
	},

	{
		"EncoderJpegGet",
		http.MethodGet,
		"/api/v1/encoder/jpeg",
		EncoderJpegGet,
	},

	{
		"EncoderJpegIdDelete",
		http.MethodDelete,
		"/api/v1/encoder/jpeg/:id",
		EncoderJpegIdDelete,
	},

	{
		"EncoderJpegIdGet",
		http.MethodGet,
		"/api/v1/encoder/jpeg/:id",
		EncoderJpegIdGet,
	},

	{
		"EncoderJpegIdPatch",
		http.MethodPatch,
		"/api/v1/encoder/jpeg/:id",
		EncoderJpegIdPatch,
	},

	{
		"EncoderJpegPost",
		http.MethodPost,
		"/api/v1/encoder/jpeg",
		EncoderJpegPost,
	},

	{
		"EncoderMjpegGet",
		http.MethodGet,
		"/api/v1/encoder/mjpeg",
		EncoderMjpegGet,
	},

	{
		"EncoderMjpegIdDelete",
		http.MethodDelete,
		"/api/v1/encoder/mjpeg/:id",
		EncoderMjpegIdDelete,
	},

	{
		"EncoderMjpegIdGet",
		http.MethodGet,
		"/api/v1/encoder/mjpeg/:id",
		EncoderMjpegIdGet,
	},

	{
		"EncoderMjpegIdPatch",
		http.MethodPatch,
		"/api/v1/encoder/mjpeg/:id",
		EncoderMjpegIdPatch,
	},

	{
		"EncoderMjpegPost",
		http.MethodPost,
		"/api/v1/encoder/mjpeg",
		EncoderMjpegPost,
	},

	{
		"StreamerGet",
		http.MethodGet,
		"/api/v1/streamer",
		StreamerGet,
	},

	{
		"StreamerJpegOverHttpGet",
		http.MethodGet,
		"/api/v1/streamer/jpeg-over-http",
		StreamerJpegOverHttpGet,
	},

	{
		"StreamerRtspGet",
		http.MethodGet,
		"/api/v1/streamer/rtsp",
		StreamerRtspGet,
	},

	{
		"StreamerVideoToStorageGet",
		http.MethodGet,
		"/api/v1/streamer/video-to-storage",
		StreamerVideoToStorageGet,
	},

	{
		"Systeminfo",
		http.MethodGet,
		"/api/v1/system",
		Systeminfo,
	},
}