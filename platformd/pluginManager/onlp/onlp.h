#ifndef ONLP_H
#define ONLP_H

#include "pluginCommon.h"

int Init();
int DeInit();

int GetPlatformState(sys_info_t *);

int GetFanState(fan_info_t *, int);
int GetAllFanState(fan_info_t *, int);
int GetMaxNumOfFans();

int GetSfpState(sfp_info_t *, int);
int GetSfpCnt();

#endif //ONPL_H
