# Network Slice Reqeust 各種 DataSet 的實驗結果 - 固定 Request's Duration - DataSet-6

###### tags: `docs` `Network Slice`

## 摘要

* Tenant 數量為 3
* 每個 Tenant 都有 10 個 Network Slice
* 固定 Network Slice Request's Duration 為 300
* Resource 對象為 CPU，上限為 1000
* Timewindow 大小為 600
* 在每個 Timewindow 中 Network Slice Reqeust 數量固定為 4 (3+1)
* DataSet 6 為 CPU lamba 為 <font color="red">1000</font> (1000 = 1 cpu)
* Forecasted discount 為 50%

## DatatSet-6

:::spoiler DataSet 產生參數
```yaml=
datasetInfo:
  name: DataSet-6          # name of the dataset
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
      lambda: 1000         # generate slice reqeusted cpu by poisson distribution
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

### DataSet-6

* [raw-data](https://github.com/p76081158/5g-nsmf/tree/assets/network-slice-reqeusts-dataset/requests/DataSet-6)

#### Sort

(1) = (2)
(5) = (6)

| Algorithm        | (1)       | (5)=(1).sort | (3)       | (7)=(3).sort | (4)    | (8)=(4).sort |
| ---------------- | --------- | ------------ | --------- | ------------ | ------ | ------------ |
| Pre-Order        | 0.5132    | 0.51115      | 0.62085   | 0.63015      | 0.6255 | 0.63565      |
| Invert-Pre-Order | 0.5111    | 0.50985      | 0.62575   | 0.63905      | 0.6361 | 0.6489       |
| Leaf-Size        | 0.5104375 | 0.50925      | 0.6126875 | 0.622625     | 0.6215 | 0.6300625    |

#### Forecasted

| Algorithm        | (1)       | (3)=(1).forecasted | (2)       | (4)=(2).forecasted | (5)     | (7)=(5).forecasted | (6)     | (8)=(6).forecasted |
| ---------------- | --------- | ------------------ | --------- | ------------------ | ------- | ------------------ | ------- | ------------------ |
| Pre-Order        | 0.5132    | 0.62085            | 0.5132    | 0.6255             | 0.51115 | 0.63015            | 0.51115 | 0.63565            |
| Invert-Pre-Order | 0.5111    | 0.62575            | 0.5111    | 0.6361             | 0.50985 | 0.63905            | 0.50985 | 0.6489             |
| Leaf-Size        | 0.5104375 | 0.6126875          | 0.5104375 | 0.6215             | 0.50925 | 0.622625           | 0.50925 | 0.6300625          |

#### Concat

(1),(2) 數值一樣
(5),(6) 數值一樣

| Algorithm        | (3)       | (4)=(3).concat | (7)      | (8)=(7).concat |
| ---------------- | --------- | -------------- | -------- | -------------- |
| Pre-Order        | 0.62085   | 0.6255         | 0.63015  | 0.63565        |
| Invert-Pre-Order | 0.62575   | 0.6361         | 0.63905  | 0.6489         |
| Leaf-Size        | 0.6126875 | 0.6215         | 0.622625 | 0.6300625      |

:::spoiler (1) result-cpu-sort-false-forecast-0-concat-false
```yaml=
datasetInfo:
  name: DataSet-6          # name of the dataset
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
      lambda: 1000          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.5132      | 0        | 948      | 51.2     | 0.8      |
| Invert-Pre-Order | 0.5111      | 0        | 955.6    | 44.4     | 0        |
| Leaf-Size        | 0.5104375   | 0        | 958.25   | 41.75    | 0        |

![](https://i.imgur.com/SlLWzGa.png)

:::spoiler (2) result-cpu-sort-false-forecast-0-concat-true <font color="red">結果同上</font>
```yaml=
datasetInfo:
  name: DataSet-6          # name of the dataset
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
      lambda: 1000          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.5132      | 0        | 948      | 51.2     | 0.8      |
| Invert-Pre-Order | 0.5111      | 0        | 955.6    | 44.4     | 0        |
| Leaf-Size        | 0.5104375   | 0        | 958.25   | 41.75    | 0        |

![](https://i.imgur.com/pRq1Kll.png)

:::spoiler (3) result-cpu-sort-false-forecast-1-concat-false
```yaml=
datasetInfo:
  name: DataSet-6          # name of the dataset
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
      lambda: 1000          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.62085     | 0        | 566.8    | 383      | 50.2     |
| Invert-Pre-Order | 0.62575     | 0        | 533.6    | 429.8    | 36.6     |
| Leaf-Size        | 0.6126875   | 0        | 589      | 371.25   | 39.75    |

![](https://i.imgur.com/Go4rCz6.png)

:::spoiler (4) result-cpu-sort-false-forecast-1-concat-true
```yaml=
datasetInfo:
  name: DataSet-6          # name of the dataset
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
      lambda: 1000          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.6255      | 0        | 549.2    | 399.6    | 51.2     |
| Invert-Pre-Order | 0.6361      | 0        | 500      | 455.6    | 44.4     |
| Leaf-Size        | 0.6215      | 0        | 560      | 394      | 46       |

![](https://i.imgur.com/WErq2UB.png)

:::spoiler (5) result-cpu-sort-true-forecast-0-concat-false
```yaml=
datasetInfo:
  name: DataSet-6          # name of the dataset
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
      lambda: 1000          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.51115     | 0        | 955.4    | 44.6     | 0        |
| Invert-Pre-Order | 0.50985     | 0        | 960.6    | 39.4     | 0        |
| Leaf-Size        | 0.50925     | 0        | 963      | 37       | 0        |

![](https://i.imgur.com/L4S1Myk.png)

:::spoiler (6) result-cpu-sort-true-forecast-0-concat-true <font color="red">結果同上</font>
```yaml=
datasetInfo:
  name: DataSet-6          # name of the dataset
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
      lambda: 1000          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.51115     | 0        | 955.4    | 44.6     | 0        |
| Invert-Pre-Order | 0.50985     | 0        | 960.6    | 39.4     | 0        |
| Leaf-Size        | 0.50925     | 0        | 963      | 37       | 0        |

![](https://i.imgur.com/L4S1Myk.png)

:::spoiler (7) result-cpu-sort-true-forecast-1-concat-false
```yaml=
datasetInfo:
  name: DataSet-6          # name of the dataset
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
      lambda: 1000          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.63015     | 0        | 524.8    | 429.8    | 45.4     |
| Invert-Pre-Order | 0.63905     | 0        | 486.2    | 471.4    | 42.4     |
| Leaf-Size        | 0.622625    | 0        | 541      | 427.5    | 31.5     |

![](https://i.imgur.com/qfFww6K.png)

:::spoiler (8) result-cpu-sort-true-forecast-1-concat-true
```yaml=
datasetInfo:
  name: DataSet-6          # name of the dataset
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
      lambda: 1000          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.63565     | 0        | 505.2    | 447      | 47.8     |
| Invert-Pre-Order | 0.6489      | 0        | 458.8    | 486.8    | 54.4     |
| Leaf-Size        | 0.6300625   | 0        | 516.25   | 447.25   | 36.5     |

![](https://i.imgur.com/C1UxgDJ.png)
