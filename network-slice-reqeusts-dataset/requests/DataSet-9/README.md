# Network Slice Reqeust 各種 DataSet 的實驗結果 - 非固定 Request's Duration - DataSet-9

###### tags: `docs` `Network Slice`

## 摘要

* Tenant 數量為 3
* 每個 Tenant 都有 10 個 Network Slice
* <font color="red">非固定</font> Network Slice Request's Duration，Duration lambda 為 300
* Resource 對象為 CPU，上限為 1000
* Timewindow 大小為 600
* 在每個 Timewindow 中 Network Slice Reqeust 數量固定為 4 (3+1)
* DataSet 9 為 CPU lamba 為 <font color="red">700</font> (500 = 0.5 cpu)
* Forecasted discount 為 50%

## DataSet-9

:::spoiler DataSet 產生參數
```yaml=
datasetInfo:
  name: DataSet-9          # name of the dataset
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
      lambda: 700          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: true           # slice request duration is fixed or random (poison lambda = duration)
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

### DataSet-9	

* [raw-data](https://github.com/p76081158/5g-nsmf/tree/assets/network-slice-reqeusts-dataset/requests/DataSet-9)

#### 總結

* 當在沒有 forecasted data 時，每個 Network Slice Reqeust 並不會被切分為幾個連續的子區塊，而且每個 Network Slice Request's Duration 是固定，所以在做 Bin Packing 擺放時不會產生碎裂問題，因此 Concatenate Algorithm 並不會有改善
* 而當有 forecasted data 時，Accept Rate 都會上升是因為 forecasted data 是由原本的 Resource 需求經過 50% discount 後產生，所以自然地 Accept Rate 會提升
* 如果再有 forecasted data 的情況下時，使用 Concatenate Algorithm 會提升 Accept Rate，其原因在於可以將碎裂的區塊串接起來利用
* Network Slice Reqeust Bin Packing 是 two-dimension bin packing problem，因此會有 X 軸與 Y 軸，由於 X 軸是有時間性的，所以 Concatenate Algorithm 是以串接水平方向的碎裂為目標
* 在使用 Concatenate Algorithm 時，Inver-Pre-Order Algorithm 提升最多 Accept Rate，原因在於 Inver-Pre-Order Algorithm 相較於其他兩個會產生較多水平方向的碎裂
* 在做 Bin Packing 前將 Reqeust 依據 Resource 大小做排序，會提升一些 Accept Rate (基本上是沒太大差異)

##### Sort

| Algorithm        | (1)      | (5)=(1).sort | (2)      | (6)=(2).sort | (3)      | (7)=(3).sort | (4)       | (8)=(4).sort |
| ---------------- | -------- | ------------ | -------- | ------------ | -------- | ------------ | --------- | ------------ |
| Pre-Order        | 0.6137   | 0.617        | 0.6139   | 0.6172       | 0.7601   | 0.75705      | 0.7678    | 0.76485      |
| Invert-Pre-Order | 0.61355  | 0.6188       | 0.61385  | 0.6198       | 0.76825  | 0.7623       | 0.7837    | 0.78155      |
| Leaf-Size        | 0.634625 | 0.634375     | 0.634875 | 0.6348125    | 0.772625 | 0.7671875    | 0.7860625 | 0.781375     |

##### Forecasted

| Algorithm        | (1)      | (3)=(1).forecasted | (2)      | (4)=(2).forecasted | (5)      | (7)=(5).forecasted | (6)       | (8)=(6).forecasted |
| ---------------- | -------- | ------------------ | -------- | ------------------ | -------- | ------------------ | --------- | ------------------ |
| Pre-Order        | 0.6137   | 0.7601             | 0.6139   | 0.7678             | 0.617    | 0.75705            | 0.6172    | 0.76485            |
| Invert-Pre-Order | 0.61355  | 0.76825            | 0.61385  | 0.7837             | 0.6188   | 0.7623             | 0.6198    | 0.78155            |
| Leaf-Size        | 0.634625 | 0.772625           | 0.634875 | 0.7860625          | 0.634375 | 0.7671875          | 0.6348125 | 0.781375           |

##### Concat

(1),(2) 數值一樣
(5),(6) 數值一樣

| Algorithm        | (1)      | (2)=(1).concat | (3)      | (4)=(3).concat | (5)      | (6)=(5).concat | (7)       | (8)=(7).concat |
| ---------------- | -------- | -------------- | -------- | -------------- | -------- | -------------- | --------- | -------------- |
| Pre-Order        | 0.6137   | 0.6139         | 0.7601   | 0.7678         | 0.617    | 0.6172         | 0.75705   | 0.76485         |
| Invert-Pre-Order | 0.61355  | 0.61385        | 0.76825  | 0.7837         | 0.6188   | 0.6198         | 0.7623    | 0.78155        |
| Leaf-Size        | 0.634625 | 0.634875       | 0.772625 | 0.7860625      | 0.634375 | 0.6348125      | 0.7671875 | 0.781375       |

:::spoiler (1) result-cpu-sort-false-forecast-0-concat-false
```yaml=
datasetInfo:
  name: DataSet-9          # name of the dataset
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
      lambda: 700          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: true           # slice request duration is fixed or random (poison lambda = duration)
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
| Pre-Order        | 0.6137      | 117.4    | 392.2    | 408.6    | 81.8     |
| Invert-Pre-Order | 0.61355     | 117.4    | 389.8    | 414      | 78.8     |
| Leaf-Size        | 0.634625    | 89       | 376.25   | 442      | 92.75    |

![](https://i.imgur.com/83DSfxo.png)

:::spoiler (2) result-cpu-sort-false-forecast-0-concat-true
```yaml=
datasetInfo:
  name: DataSet-9          # name of the dataset
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
      lambda: 700          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: true           # slice request duration is fixed or random (poison lambda = duration)
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
| Pre-Order        | 0.6139      | 117.4    | 391.8    | 408.6    | 82.2     |
| Invert-Pre-Order | 0.61385     | 117.4    | 389.2    | 414      | 79.4     |
| Leaf-Size        | 0.634875    | 89       | 375.75   | 442      | 93.25    |

![](https://i.imgur.com/hQ3L3Fy.png)

:::spoiler (3) result-cpu-sort-false-forecast-1-concat-false
```yaml=
datasetInfo:
  name: DataSet-9          # name of the dataset
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
      lambda: 700          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: true           # slice request duration is fixed or random (poison lambda = duration)
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
| Pre-Order        | 0.7601      | 12.4     | 206.4    | 509.6    | 271.6    |
| Invert-Pre-Order | 0.76825     | 12.4     | 193.4    | 503      | 291.2    |
| Leaf-Size        | 0.772625    | 10.5     | 185      | 508      | 296.5    |

![](https://i.imgur.com/wUXpCJq.png)

:::spoiler (4) result-cpu-sort-false-forecast-1-concat-true
```yaml=
datasetInfo:
  name: DataSet-9          # name of the dataset
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
      lambda: 700          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: true           # slice request duration is fixed or random (poison lambda = duration)
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
| Pre-Order        | 0.7678      | 9.6      | 193.6    | 512.8    | 284      |
| Invert-Pre-Order | 0.7837      | 9.6      | 172.8    | 490.8    | 326.8    |
| Leaf-Size        | 0.7860625   | 7        | 162.5    | 509.75   | 320.75   |

![](https://i.imgur.com/SYTZrkI.png)

:::spoiler (5) result-cpu-sort-true-forecast-0-concat-false
```yaml=
datasetInfo:
  name: DataSet-9          # name of the dataset
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
      lambda: 700          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: true           # slice request duration is fixed or random (poison lambda = duration)
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
| Pre-Order        | 0.617       | 84.8     | 437      | 403.6    | 74.6     |
| Invert-Pre-Order | 0.6188      | 84.8     | 434.6    | 401.2    | 79.4     |
| Leaf-Size        | 0.634375    | 72.25    | 408.75   | 428.25   | 90.75    |

![](https://i.imgur.com/CisvuPE.png)

:::spoiler (6) result-cpu-sort-true-forecast-0-concat-true
```yaml=
datasetInfo:
  name: DataSet-9          # name of the dataset
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
      lambda: 700          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: true           # slice request duration is fixed or random (poison lambda = duration)
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
| Pre-Order        | 0.6172      | 84.8     | 436.4    | 404      | 74.8     |
| Invert-Pre-Order | 0.6198      | 84.8     | 432.2    | 402      | 81       |
| Leaf-Size        | 0.6348125   | 72.25    | 407.5    | 429      | 91.25    |

![](https://i.imgur.com/VusNTiB.png)

:::spoiler (7) result-cpu-sort-true-forecast-1-concat-false
```yaml=
datasetInfo:
  name: DataSet-9          # name of the dataset
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
      lambda: 700          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: true           # slice request duration is fixed or random (poison lambda = duration)
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
| Pre-Order        | 0.75705     | 9.6      | 212.6    | 517.8    | 260      |
| Invert-Pre-Order | 0.7623      | 9.6      | 206.6    | 508.8    | 275      |
| Leaf-Size        | 0.7671875   | 7.5      | 199.5    | 509.75   | 283.25   |

![](https://i.imgur.com/LlCKTnz.png)

:::spoiler (8) result-cpu-sort-true-forecast-1-concat-true
```yaml=
datasetInfo:
  name: DataSet-9          # name of the dataset
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
      lambda: 700          # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    bandwidth:
      limit: 10            # bandwidth reqeust limit
      lambda: 5            # generate slice reqeusted cpu by poisson distribution
      discount: 0.5        # forecasting discount
    duration: 300          # duration of slice reqeust
    random: true           # slice request duration is fixed or random (poison lambda = duration)
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
| Pre-Order        | 0.76485     | 7.2      | 197.6    | 523.8    | 271.4    |
| Invert-Pre-Order | 0.78155     | 7.2      | 175.8    | 500.6    | 316.4    |
| Leaf-Size        | 0.781375    | 5.25     | 174.75   | 509.25   | 310.75   |

![](https://i.imgur.com/yGpQwgQ.png)