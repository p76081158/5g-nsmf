# Network Slice Reqeust 各種 DataSet 的實驗結果 - 固定 Request's Duration - DataSet-4

###### tags: `docs` `Network Slice`

## 摘要

* Tenant 數量為 3
* 每個 Tenant 都有 10 個 Network Slice
* 固定 Network Slice Request's Duration 為 300
* Resource 對象為 CPU，上限為 1000
* Timewindow 大小為 600
* 在每個 Timewindow 中 Network Slice Reqeust 數量固定為 4 (3+1)
* DataSet 4 為 CPU lamba 為 <font color="red">800</font> (800 = 0.8 cpu)
* Forecasted discount 為 50%

## DatatSet-4

:::spoiler DataSet 產生參數
```yaml=
datasetInfo:
  name: DataSet-4          # name of the dataset
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
      lambda: 800          # generate slice reqeusted cpu by poisson distribution
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

### DataSet-4

* [raw-data](https://github.com/p76081158/5g-nsmf/tree/assets/network-slice-reqeusts-dataset/requests/DataSet-4)

#### Sort

(1) = (2)
(5) = (6)

| Algorithm        | (1)     | (5)=(1).sort | (3)       | (7)=(3).sort | (4)    | (8)=(4).sort |
| ---------------- | ------- | ------------ | --------- | ------------ | ------ | ------------ |
| Pre-Order        | 0.57    | 0.5707       | 0.71285   | 0.70455      | 0.7278 | 0.721        |
| Invert-Pre-Order | 0.5635  | 0.56525      | 0.71855   | 0.71965      | 0.7525 | 0.74735      |
| Leaf-Size        | 0.56625 | 0.5664375    | 0.7108125 | 0.712375     | 0.7355 | 0.737625     |

#### Forecasted

| Algorithm        | (1)     | (3)=(1).forecasted | (2)     | (4)=(2).forecasted | (5)       | (7)=(5).forecasted | (6)       | (8)=(6).forecasted |
| ---------------- | ------- | ------------------ | ------- | ------------------ | --------- | ------------------ | --------- | ------------------ |
| Pre-Order        | 0.57    | 0.71285            | 0.57    | 0.7278             | 0.5707    | 0.70455            | 0.5707    | 0.721              |
| Invert-Pre-Order | 0.5635  | 0.71855            | 0.5635  | 0.7525             | 0.56525   | 0.71965            | 0.56525   | 0.74735            |
| Leaf-Size        | 0.56625 | 0.7108125          | 0.56625 | 0.7355             | 0.5664375 | 0.712375           | 0.5664375 | 0.737625           |

#### Concat

(1),(2) 數值一樣
(5),(6) 數值一樣

| Algorithm        | (3)       | (4)=(3).concat | (7)      | (8)=(7).concat |
| ---------------- | --------- | -------------- | -------- | -------------- |
| Pre-Order        | 0.71285   | 0.7278         | 0.70455  | 0.721          |
| Invert-Pre-Order | 0.71855   | 0.7525         | 0.71965  | 0.74735        |
| Leaf-Size        | 0.7108125 | 0.7355         | 0.712375 | 0.737625       |

:::spoiler (1) result-cpu-sort-false-forecast-0-concat-false
```yaml=
datasetInfo:
  name: DataSet-4          # name of the dataset
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
      lambda: 800          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.57        | 0        | 747.4    | 225.2    | 27.4     |
| Invert-Pre-Order | 0.5635      | 0        | 762.8    | 220.4    | 16.8     |
| Leaf-Size        | 0.56625     | 0        | 770.25   | 194.5    | 35.25    |

![](https://i.imgur.com/RhFXvhI.png)

:::spoiler (2) result-cpu-sort-false-forecast-0-concat-true <font color="red">結果同上</font>
```yaml=
datasetInfo:
  name: DataSet-4          # name of the dataset
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
      lambda: 800          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.57        | 0        | 747.4    | 225.2    | 27.4     |
| Invert-Pre-Order | 0.5635      | 0        | 762.8    | 220.4    | 16.8     |
| Leaf-Size        | 0.56625     | 0        | 770.25   | 194.5    | 35.25    |

![](https://i.imgur.com/RhFXvhI.png)

:::spoiler (3) result-cpu-sort-false-forecast-1-concat-false
```yaml=
datasetInfo:
  name: DataSet-4          # name of the dataset
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
      lambda: 800          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.71285     | 0        | 304.4    | 539.8    | 155.8    |
| Invert-Pre-Order | 0.71855     | 0        | 267.4    | 591      | 141.6    |
| Leaf-Size        | 0.7108125   | 0        | 299.75   | 557.25   | 143      |

![](https://i.imgur.com/suT1JdP.png)

:::spoiler (4) result-cpu-sort-false-forecast-1-concat-true
```yaml=
datasetInfo:
  name: DataSet-4          # name of the dataset
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
      lambda: 800          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.7278      | 0        | 258.8    | 571.2    | 170      |
| Invert-Pre-Order | 0.7525      | 0        | 181.2    | 627.6    | 191.2    |
| Leaf-Size        | 0.7355      | 0        | 231.5    | 595      | 173.5    |

![](https://i.imgur.com/WUK9aT2.png)

:::spoiler (5) result-cpu-sort-true-forecast-0-concat-false
```yaml=
datasetInfo:
  name: DataSet-4          # name of the dataset
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
      lambda: 800          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.5707      | 0        | 742      | 233.2    | 24.8     |
| Invert-Pre-Order | 0.56525     | 0        | 755.2    | 228.6    | 16.2     |
| Leaf-Size        | 0.5664375   | 0        | 768.25   | 197.75   | 34       |

![](https://i.imgur.com/7xkEQ5r.png)

:::spoiler (6) result-cpu-sort-true-forecast-0-concat-true <font color="red">結果同上</font>
```yaml=
datasetInfo:
  name: DataSet-4          # name of the dataset
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
      lambda: 800          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.5707      | 0        | 742      | 233.2    | 24.8     |
| Invert-Pre-Order | 0.56525     | 0        | 755.2    | 228.6    | 16.2     |
| Leaf-Size        | 0.5664375   | 0        | 768.25   | 197.75   | 34       |

![](https://i.imgur.com/7xkEQ5r.png)

:::spoiler (7) result-cpu-sort-true-forecast-1-concat-false
```yaml=
datasetInfo:
  name: DataSet-4          # name of the dataset
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
      lambda: 800          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.70455     | 0        | 333.2    | 515.4    | 151.4    |
| Invert-Pre-Order | 0.71965     | 0        | 269.4    | 582.6    | 148      |
| Leaf-Size        | 0.712375    | 0        | 304.5    | 541.5    | 154      |

![](https://i.imgur.com/7uvZ8UP.png)

:::spoiler (8) result-cpu-sort-true-forecast-1-concat-true
```yaml=
datasetInfo:
  name: DataSet-4          # name of the dataset
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
      lambda: 800          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.721       | 0        | 281.2    | 553.6    | 165.2    |
| Invert-Pre-Order | 0.74735     | 0        | 201.8    | 607      | 191.2    |
| Leaf-Size        | 0.737625    | 0        | 234      | 581.5    | 184.5    |

![](https://i.imgur.com/13inZBV.png)
