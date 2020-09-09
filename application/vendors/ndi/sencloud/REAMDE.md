```
Hi Nikita…below is your Vendor ID…if you have any technical questions about this you can respond directly to Brian but please cc me so I can help support you both 😊.

 

Regards, Michael

 

From: Brian Brice <bbrice@newtek.com>
Sent: Tuesday, September 8, 2020 2:17 PM
To: Michael Kornet <mkornet@ndi.tv>

Subject: RE: Vendor ID

 

Hello.

 

I have uploaded a new build of the NDI Embedded SDK that incorporates your new vendor ID here:

https://1drv.ms/u/s!ArRUZ4qdf_bVoStLV-httrkqEDVO

 

Here is the new vendor information:

 

    Vendor name: Anhui SEN-Cloud

    Vendor ID: E4651A8F-C223-472E-9467-D9756B2A6C27

 

The following is an example of how to use the vendor information:

 

    const char* ndi_config = R"({

      "ndi": {

        "vendor": {

          "name": "Anhui SEN-Cloud",

          "id": "E4651A8F-C223-472E-9467-D9756B2A6C27"

        }

      }

    })";

    NDIlib_send_instance_t ndi_sender = NDIlib_send_create_v2(nullptr, ndi_config);

 

Thank you!

 

--

Brian Brice
```
