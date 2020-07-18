#if !defined(__IMX290_DEF_H_)
#define __IMX290_DEF_H_

#define SENSOR_IMX290_1080P_30FPS_MODE      (0)
#define SENSOR_IMX290_1080P_60FPS_MODE      (1)

/*
NOTE:
exposure time is limited when use 10bit 720P WDR mode 60fps,
and the exposure time can reach to its max at 56.70fps or less
*/
#define SENSOR_IMX290_720P_60FPS_MODE       (2)
#define SENSOR_IMX290_720P_120FPS_MODE      (3)


/*
0:10bit 1:12bit 
NOTE:
exposure time is limited when use 12bit mode 30fps, and 
the exposure time can reach to its max at 28.36fps or less

10bit mode's exposure time is OK at 30fps.
*/
#define SENSOR_IMX290_LINE_WDR_12BIT      (0)



#endif 

