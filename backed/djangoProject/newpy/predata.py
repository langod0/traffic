import os
import pandas as pd

from rest_framework import viewsets
from rest_framework.decorators import action
from rest_framework.response import Response

class predata(viewsets.ViewSet):
    @action(methods=['post'],detail=False)
    def pre(self,request):
        test_path="newpy/data/submission29_1.csv"
        tmp = pd.read_csv(os.path.join(test_path))
        tmp = self.getnum1(tmp)
        numin=tmp[0]
        numout=tmp[1]
        station_num_in=tmp[2]
        station_num_out=tmp[3]
        tm=[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23]
        stationnum=81
        result={
            "Numin":numin,
            "Numout":numout,
            "station_num_in":station_num_in,
            "station_num_out":station_num_out,
            "Tm":tm,
            "Stationnum":stationnum
        }
        return Response(result)


    def getnum1(self,df_):
        train_df = df_.copy()
        numin = []
        numout=[]
        station_in=[]
        station_out=[]
        for i in range(0, 24):
            numin.append(0)
            numout.append(0)
        for j in range(0,81):
            t=[]
            for i in range(0,24):
                t.append(0)
            station_in.append(t)
            station_out.append(t)
        hour = train_df['startTime'].apply(lambda x: int(x[11:13]))
        # hour2 = train_df['endTime'].apply(lambda x: int(x[11:13]))
        numi = train_df['inNums']
        numo= train_df['outNums']
        station = train_df['stationID']
        for i in range(0,len(numi)):
            if numi[i]>0:
                numin[hour[i]] += numi[i]
                station_in[station[i]][hour[i]] += numi[i]
            if numo[i]>0:
                numout[hour[i]] += numo[i]
                station_out[station[i]][hour[i]] += numo[i]
        return [numin,numout,station_in,station_out]
