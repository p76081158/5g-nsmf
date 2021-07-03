# Network Slice Reqeust 各種 DataSet 的實驗結果 - 固定 Request's Duration - DataSet-2

###### tags: `docs` `Network Slice`

## 摘要

* Tenant 數量為 3
* 每個 Tenant 都有 10 個 Network Slice
* 固定 Network Slice Request's Duration 為 300
* Resource 對象為 CPU，上限為 1000
* Timewindow 大小為 600
* 在每個 Timewindow 中 Network Slice Reqeust 數量固定為 4 (3+1)
* DataSet 2 為 CPU lamba 為 <font color="red">600</font> (600 = 0.6 cpu)
* Forecasted discount 為 50%

## DataSet-2

:::spoiler DataSet 產生參數
```yaml=
datasetInfo:
  name: DataSet-2          # name of the dataset
  ngciList:                # tenants list
    - 466-01-000000010
    - 466-11-000000010
    - 466-93-000000010
  sliceNum: 10             # slice number of each tenants
  extraRequest: 1          # each timewindow has n network slice requests ( n = tenant num + extraRequest, e.g. n = 3 + 1 )
  testNum: 5               # how many test case in this dataset
  resource:
    cpu:
      limit: 1000          # cpu request limit  (1000 = 1 cpu core)
      lambda: 600          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: false          # slice request duration is fixed or random (poison lambda = duration)
  target: cpu              # target of resource for slice bin packing algorithm
  sort: false              # sorting network slice requests in each timewindow before bin packing
  concat: false            # use concat algorithm if not finding any tree node for slice reqeust
  timewindow: 600          # duration of each timewindow
  forecastingTime: 0       # from which timewindow start use forecasted data (0 == nerver, 1 == from timewindow-1)
  forecastBlockSize: 100   # forecasted block size of network slice request, e.g. each network slice reqeust's duration is 300, so each slice requests will be divided into two sub slice requests (300/150=2)
  regenerate: true         # regenerate DataSet again (false for DataSet already exist)
```
:::

## 實驗結果

### 實驗說明

result-<1>-sort-<2>-forecast-<3>-concat-<4>

1. Bin Packing 的對象
2. Bin Packing 前是否排序
3. 從第幾個 timewindow 開始使用 forecasted data (0 代表不使用)
4. 當 Bin Packing 演算法找不到位置放 Network Slice Reqeust 時，是否使用 Concatenate 演算法

### DataSet-2	

* [raw-data](https://github.com/p76081158/5g-nsmf/tree/assets/network-slice-reqeusts-dataset/requests/DataSet-2)

#### Sort

(1) = (2)
(5) = (6)

| Algorithm        | (1)    | (5)=(1).sort | (3)      | (7)=(3).sort | (4)     | (8)=(4).sort |
| ---------------- | ------ | ------------ | -------- | ------------ | ------- | ------------ |
| Pre-Order        | 0.728  | 0.73315      | 0.80595  | 0.8123       | 0.82315 | 0.8329       |
| Invert-Pre-Order | 0.7129 | 0.71475      | 0.81295  | 0.81085      | 0.8592  | 0.85745      |
| Leaf-Size        | 0.7372 | 0.731875     | 0.799125 | 0.7975625    | 0.8338  | 0.8358125    |

#### Forecasted

| Algorithm        | (1)    | (3)=(1).forecasted | (2)    | (4)=(2).forecasted | (5)      | (7)=(5).forecasted | (6)      | (8)=(6).forecasted |
| ---------------- | ------ | ------------------ | ------ | ------------------ | -------- | ------------------ | -------- | ------------------ |
| Pre-Order        | 0.728  | 0.80595            | 0.728  | 0.82315            | 0.73315  | 0.8123             | 0.73315  | 0.8329             |
| Invert-Pre-Order | 0.7129 | 0.81295            | 0.7129 | 0.8592             | 0.71475  | 0.81085            | 0.71475  | 0.85745            |
| Leaf-Size        | 0.7372 | 0.799125           | 0.7372 | 0.8338             | 0.731875 | 0.7975625          | 0.731875 | 0.8358125          |

#### Concat

(1),(2) 數值一樣
(5),(6) 數值一樣

| Algorithm        | (3)      | (4)=(3).concat | (7)       | (8)=(7).concat |
| ---------------- | -------- | -------------- | --------- | -------------- |
| Pre-Order        | 0.80595  | 0.82315        | 0.8123    | 0.8329         |
| Invert-Pre-Order | 0.81295  | 0.8592         | 0.81085   | 0.85745        |
| Leaf-Size        | 0.799125 | 0.8338         | 0.7975625 | 0.8358125      |


:::spoiler (1) result-cpu-sort-false-forecast-0-concat-false
```yaml=
datasetInfo:
  name: DataSet-2          # name of the dataset
  ngciList:                # tenants list
    - 466-01-000000010
    - 466-11-000000010
    - 466-93-000000010
  sliceNum: 10             # slice number of each tenants
  extraRequest: 1          # each timewindow has n network slice requests ( n = tenant num + extraRequest, e.g. n = 3 + 1 )
  testNum: 5               # how many test case in this dataset
  resource:
    cpu:
      limit: 1000          # cpu request limit  (1000 = 1 cpu core)
      lambda: 600          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: false          # slice request duration is fixed or random (poison lambda = duration)
  target: cpu              # target of resource for slice bin packing algorithm
  sort: false              # sorting network slice requests in each timewindow before bin packing
  concat: false            # use concat algorithm if not finding any tree node for slice reqeust
  timewindow: 600          # duration of each timewindow
  forecastingTime: 0       # from which timewindow start use forecasted data (0 == nerver, 1 == from timewindow-1)
  forecastBlockSize: 100   # forecasted block size of network slice request, e.g. each network slice reqeust's duration is 300, so each slice requests will be divided into two sub slice requests (300/150=2)
  regenerate: false        # regenerate DataSet again (false for DataSet already exist)
```
:::

| Algorithm        | Accept Rate | Accept-1 | Accept-2 | Accept-3 | Accept-4 |
| ---------------- | ----------- | -------- | -------- | -------- | -------- |
| Pre-Order        | 0.728       | 0        | 272.6    | 542.8    | 184.6    |
| Invert-Pre-Order | 0.7129      | 0        | 301.8    | 544.8    | 153.4    |
| Leaf-Size        | 0.7372      | 0        | 275.8    | 499.6    | 224.6    |

![](https://i.imgur.com/P5tgE1I.png)

:::spoiler (2) result-cpu-sort-false-forecast-0-concat-true <font color="red">結果同上</font>
```yaml=
datasetInfo:
  name: DataSet-2          # name of the dataset
  ngciList:                # tenants list
    - 466-01-000000010
    - 466-11-000000010
    - 466-93-000000010
  sliceNum: 10             # slice number of each tenants
  extraRequest: 1          # each timewindow has n network slice requests ( n = tenant num + extraRequest, e.g. n = 3 + 1 )
  testNum: 5               # how many test case in this dataset
  resource:
    cpu:
      limit: 1000          # cpu request limit  (1000 = 1 cpu core)
      lambda: 600          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: false          # slice request duration is fixed or random (poison lambda = duration)
  target: cpu              # target of resource for slice bin packing algorithm
  sort: false              # sorting network slice requests in each timewindow before bin packing
  concat: true             # use concat algorithm if not finding any tree node for slice reqeust
  timewindow: 600          # duration of each timewindow
  forecastingTime: 0       # from which timewindow start use forecasted data (0 == nerver, 1 == from timewindow-1)
  forecastBlockSize: 100   # forecasted block size of network slice request, e.g. each network slice reqeust's duration is 300, so each slice requests will be divided into two sub slice requests (300/150=2)
  regenerate: false        # regenerate DataSet again (false for DataSet already exist)
```
:::

| Algorithm        | Accept Rate | Accept-1 | Accept-2 | Accept-3 | Accept-4 |
| ---------------- | ----------- | -------- | -------- | -------- | -------- |
| Pre-Order        | 0.728       | 0        | 272.6    | 542.8    | 184.6    |
| Invert-Pre-Order | 0.7129      | 0        | 301.8    | 544.8    | 153.4    |
| Leaf-Size        | 0.7372      | 0        | 275.8    | 499.6    | 224.6    |

![](https://i.imgur.com/P5tgE1I.png)

:::spoiler (3) result-cpu-sort-false-forecast-1-concat-false
```yaml=
datasetInfo:
  name: DataSet-2          # name of the dataset
  ngciList:                # tenants list
    - 466-01-000000010
    - 466-11-000000010
    - 466-93-000000010
  sliceNum: 10             # slice number of each tenants
  extraRequest: 1          # each timewindow has n network slice requests ( n = tenant num + extraRequest, e.g. n = 3 + 1 )
  testNum: 5               # how many test case in this dataset
  resource:
    cpu:
      limit: 1000          # cpu request limit  (1000 = 1 cpu core)
      lambda: 600          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: false          # slice request duration is fixed or random (poison lambda = duration)
  target: cpu              # target of resource for slice bin packing algorithm
  sort: false              # sorting network slice requests in each timewindow before bin packing
  concat: false            # use concat algorithm if not finding any tree node for slice reqeust
  timewindow: 600          # duration of each timewindow
  forecastingTime: 1       # from which timewindow start use forecasted data (0 == nerver, 1 == from timewindow-1)
  forecastBlockSize: 100   # forecasted block size of network slice request, e.g. each network slice reqeust's duration is 300, so each slice requests will be divided into two sub slice requests (300/150=2)
  regenerate: false        # regenerate DataSet again (false for DataSet already exist)
```
:::

| Algorithm        | Accept Rate | Accept-1 | Accept-2 | Accept-3 | Accept-4 |
| ---------------- | ----------- | -------- | -------- | -------- | -------- |
| Pre-Order        | 0.80595     | 0        | 161.4    | 453.4    | 385.2    |
| Invert-Pre-Order | 0.81295     | 0        | 123.6    | 501      | 375.4    |
| Leaf-Size        | 0.799125    | 0        | 169      | 465.5    | 365.5    |

![](https://i.imgur.com/yT9ESmX.png)

:::spoiler (4) result-cpu-sort-false-forecast-1-concat-true
```yaml=
datasetInfo:
  name: DataSet-2          # name of the dataset
  ngciList:                # tenants list
    - 466-01-000000010
    - 466-11-000000010
    - 466-93-000000010
  sliceNum: 10             # slice number of each tenants
  extraRequest: 1          # each timewindow has n network slice requests ( n = tenant num + extraRequest, e.g. n = 3 + 1 )
  testNum: 5               # how many test case in this dataset
  resource:
    cpu:
      limit: 1000          # cpu request limit  (1000 = 1 cpu core)
      lambda: 600          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: false          # slice request duration is fixed or random (poison lambda = duration)
  target: cpu              # target of resource for slice bin packing algorithm
  sort: false              # sorting network slice requests in each timewindow before bin packing
  concat: true             # use concat algorithm if not finding any tree node for slice reqeust
  timewindow: 600          # duration of each timewindow
  forecastingTime: 1       # from which timewindow start use forecasted data (0 == nerver, 1 == from timewindow-1)
  forecastBlockSize: 100   # forecasted block size of network slice request, e.g. each network slice reqeust's duration is 300, so each slice requests will be divided into two sub slice requests (300/150=2)
  regenerate: false        # regenerate DataSet again (false for DataSet already exist)
```
:::

| Algorithm        | Accept Rate | Accept-1 | Accept-2 | Accept-3 | Accept-4 |
| ---------------- | ----------- | -------- | -------- | -------- | -------- |
| Pre-Order        | 0.82315     | 0        | 120      | 467.4    | 412.6    |
| Invert-Pre-Order | 0.8592      | 0        | 60.8     | 441.6    | 497.6    |
| Leaf-Size        | 0.8338      | 0        | 102.4    | 460      | 437.6    |

![](https://i.imgur.com/V4uO349.png)

:::spoiler (5) result-cpu-sort-true-forecast-0-concat-false
```yaml=
datasetInfo:
  name: DataSet-2          # name of the dataset
  ngciList:                # tenants list
    - 466-01-000000010
    - 466-11-000000010
    - 466-93-000000010
  sliceNum: 10             # slice number of each tenants
  extraRequest: 1          # each timewindow has n network slice requests ( n = tenant num + extraRequest, e.g. n = 3 + 1 )
  testNum: 5               # how many test case in this dataset
  resource:
    cpu:
      limit: 1000          # cpu request limit  (1000 = 1 cpu core)
      lambda: 600          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: false          # slice request duration is fixed or random (poison lambda = duration)
  target: cpu              # target of resource for slice bin packing algorithm
  sort: true               # sorting network slice requests in each timewindow before bin packing
  concat: false            # use concat algorithm if not finding any tree node for slice reqeust
  timewindow: 600          # duration of each timewindow
  forecastingTime: 0       # from which timewindow start use forecasted data (0 == nerver, 1 == from timewindow-1)
  forecastBlockSize: 100   # forecasted block size of network slice request, e.g. each network slice reqeust's duration is 300, so each slice requests will be divided into two sub slice requests (300/150=2)
  regenerate: false        # regenerate DataSet again (false for DataSet already exist)
```
:::

| Algorithm        | Accept Rate | Accept-1 | Accept-2 | Accept-3 | Accept-4 |
| ---------------- | ----------- | -------- | -------- | -------- | -------- |
| Pre-Order        | 0.73315     | 0        | 273.4    | 520.6    | 206      |
| Invert-Pre-Order | 0.71475     | 0        | 303.2    | 534.6    | 162.2    |
| Leaf-Size        | 0.731875    | 0        | 298.75   | 475      | 226.25   |

![](https://i.imgur.com/5Y8Khvv.png)

:::spoiler (6) result-cpu-sort-true-forecast-0-concat-true <font color="red">結果同上</font>
```yaml=
datasetInfo:
  name: DataSet-2          # name of the dataset
  ngciList:                # tenants list
    - 466-01-000000010
    - 466-11-000000010
    - 466-93-000000010
  sliceNum: 10             # slice number of each tenants
  extraRequest: 1          # each timewindow has n network slice requests ( n = tenant num + extraRequest, e.g. n = 3 + 1 )
  testNum: 5               # how many test case in this dataset
  resource:
    cpu:
      limit: 1000          # cpu request limit  (1000 = 1 cpu core)
      lambda: 600          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: false          # slice request duration is fixed or random (poison lambda = duration)
  target: cpu              # target of resource for slice bin packing algorithm
  sort: true               # sorting network slice requests in each timewindow before bin packing
  concat: true             # use concat algorithm if not finding any tree node for slice reqeust
  timewindow: 600          # duration of each timewindow
  forecastingTime: 0       # from which timewindow start use forecasted data (0 == nerver, 1 == from timewindow-1)
  forecastBlockSize: 100   # forecasted block size of network slice request, e.g. each network slice reqeust's duration is 300, so each slice requests will be divided into two sub slice requests (300/150=2)
  regenerate: false        # regenerate DataSet again (false for DataSet already exist)
```
:::

| Algorithm        | Accept Rate | Accept-1 | Accept-2 | Accept-3 | Accept-4 |
| ---------------- | ----------- | -------- | -------- | -------- | -------- |
| Pre-Order        | 0.73315     | 0        | 273.4    | 520.6    | 206      |
| Invert-Pre-Order | 0.71475     | 0        | 303.2    | 534.6    | 162.2    |
| Leaf-Size        | 0.731875    | 0        | 298.75   | 475      | 226.25   |

![](https://i.imgur.com/5Y8Khvv.png)

:::spoiler (7) result-cpu-sort-true-forecast-1-concat-false
```yaml=
datasetInfo:
  name: DataSet-2          # name of the dataset
  ngciList:                # tenants list
    - 466-01-000000010
    - 466-11-000000010
    - 466-93-000000010
  sliceNum: 10             # slice number of each tenants
  extraRequest: 1          # each timewindow has n network slice requests ( n = tenant num + extraRequest, e.g. n = 3 + 1 )
  testNum: 5               # how many test case in this dataset
  resource:
    cpu:
      limit: 1000          # cpu request limit  (1000 = 1 cpu core)
      lambda: 600          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: false          # slice request duration is fixed or random (poison lambda = duration)
  target: cpu              # target of resource for slice bin packing algorithm
  sort: true               # sorting network slice requests in each timewindow before bin packing
  concat: false            # use concat algorithm if not finding any tree node for slice reqeust
  timewindow: 600          # duration of each timewindow
  forecastingTime: 1       # from which timewindow start use forecasted data (0 == nerver, 1 == from timewindow-1)
  forecastBlockSize: 100   # forecasted block size of network slice request, e.g. each network slice reqeust's duration is 300, so each slice requests will be divided into two sub slice requests (300/150=2)
  regenerate: false        # regenerate DataSet again (false for DataSet already exist)
```
:::

| Algorithm        | Accept Rate | Accept-1 | Accept-2 | Accept-3 | Accept-4 |
| ---------------- | ----------- | -------- | -------- | -------- | -------- |
| Pre-Order        | 0.8123      | 0        | 147.2    | 456.4    | 396.4    |
| Invert-Pre-Order | 0.81085     | 0        | 115.2    | 526.2    | 358.6    |
| Leaf-Size        | 0.7975625   | 0        | 167.5    | 474.75   | 357.75   |

![](https://i.imgur.com/SSzIuqT.png)

:::spoiler (8) result-cpu-sort-true-forecast-1-concat-true
```yaml=
datasetInfo:
  name: DataSet-2          # name of the dataset
  ngciList:                # tenants list
    - 466-01-000000010
    - 466-11-000000010
    - 466-93-000000010
  sliceNum: 10             # slice number of each tenants
  extraRequest: 1          # each timewindow has n network slice requests ( n = tenant num + extraRequest, e.g. n = 3 + 1 )
  testNum: 5               # how many test case in this dataset
  resource:
    cpu:
      limit: 1000          # cpu request limit  (1000 = 1 cpu core)
      lambda: 600          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: false          # slice request duration is fixed or random (poison lambda = duration)
  target: cpu              # target of resource for slice bin packing algorithm
  sort: true               # sorting network slice requests in each timewindow before bin packing
  concat: true             # use concat algorithm if not finding any tree node for slice reqeust
  timewindow: 600          # duration of each timewindow
  forecastingTime: 1       # from which timewindow start use forecasted data (0 == nerver, 1 == from timewindow-1)
  forecastBlockSize: 100   # forecasted block size of network slice request, e.g. each network slice reqeust's duration is 300, so each slice requests will be divided into two sub slice requests (300/150=2)
  regenerate: false        # regenerate DataSet again (false for DataSet already exist)
```
:::

| Algorithm        | Accept Rate | Accept-1 | Accept-2 | Accept-3 | Accept-4 |
| ---------------- | ----------- | -------- | -------- | -------- | -------- |
| Pre-Order        | 0.8329      | 0        | 109      | 450.4    | 440.6    |
| Invert-Pre-Order | 0.85745     | 0        | 61.6     | 447      | 491.4    |
| Leaf-Size        | 0.8358125   | 0        | 103      | 450.75   | 446.25   |

![](https://i.imgur.com/zKFuq1D.png)