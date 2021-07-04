# Network Slice Reqeust 各種 DataSet 的實驗結果 - 非固定 Request's Duration - DataSet-11

###### tags: `docs` `Network Slice`

## 摘要

* Tenant 數量為 3
* 每個 Tenant 都有 10 個 Network Slice
* <font color="red">非固定</font> Network Slice Request's Duration，Duration lambda 為 300
* Resource 對象為 CPU，上限為 1000
* Timewindow 大小為 600
* 在每個 Timewindow 中 Network Slice Reqeust 數量固定為 4 (3+1)
* DataSet 11 為 CPU lamba 為 <font color="red">900</font> (500 = 0.5 cpu)
* Forecasted discount 為 50%

## DataSet-11

:::spoiler DataSet 產生參數
```yaml=
datasetInfo:
  name: DataSet-11         # name of the dataset
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
      lambda: 900          # generate slice reqeusted cpu by poisson distribution
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

### DataSet-11

* [raw-data](https://github.com/p76081158/5g-nsmf/tree/assets/network-slice-reqeusts-dataset/requests/DataSet-11)

#### 總結

* 當在沒有 forecasted data 時，每個 Network Slice Reqeust 並不會被切分為幾個連續的子區塊，而且每個 Network Slice Request's Duration 是固定，所以在做 Bin Packing 擺放時不會產生碎裂問題，因此 Concatenate Algorithm 並不會有改善
* 而當有 forecasted data 時，Accept Rate 都會上升是因為 forecasted data 是由原本的 Resource 需求經過 50% discount 後產生，所以自然地 Accept Rate 會提升
* 如果再有 forecasted data 的情況下時，使用 Concatenate Algorithm 會提升 Accept Rate，其原因在於可以將碎裂的區塊串接起來利用
* Network Slice Reqeust Bin Packing 是 two-dimension bin packing problem，因此會有 X 軸與 Y 軸，由於 X 軸是有時間性的，所以 Concatenate Algorithm 是以串接水平方向的碎裂為目標
* 在使用 Concatenate Algorithm 時，Inver-Pre-Order Algorithm 提升最多 Accept Rate，原因在於 Inver-Pre-Order Algorithm 相較於其他兩個會產生較多水平方向的碎裂
* 在做 Bin Packing 前將 Reqeust 依據 Resource 大小做排序，會提升一些 Accept Rate (基本上是沒太大差異)

##### Sort

| Algorithm        | (1)       | (5)=(1).sort | (2)      | (6)=(2).sort | (3)       | (7)=(3).sort | (4)       | (8)=(4).sort |
| ---------------- | --------- | ------------ | -------- | ------------ | --------- | ------------ | --------- | ------------ |
| Pre-Order        | 0.54575   | 0.5575       | 0.5458   | 0.5576      | 0.69665   | 0.68905       | 0.69985   | 0.6936       |
| Invert-Pre-Order | 0.54435   | 0.55655      | 0.5444   | 0.55675       | 0.69105   | 0.68965       | 0.6968    | 0.6992       |
| Leaf-Size        | 0.5565625 | 0.5585       | 0.556625 | 0.5585    | 0.6869375 | 0.6938125     | 0.6909375 | 0.7005625      |

##### Forecasted

| Algorithm        | (1)       | (3)=(1).forecasted | (2)      | (4)=(2).forecasted | (5)     | (7)=(5).forecasted | (6)       | (8)=(6).forecasted |
| ---------------- | --------- | ------------------ | -------- | ------------------ | ------- | ------------------ | --------- | ------------------ |
| Pre-Order        | 0.54575   | 0.69665            | 0.5458   | 0.69985            | 0.5575  | 0.68905             | 0.5576   | 0.6936             |
| Invert-Pre-Order | 0.54435   | 0.69105            | 0.5444   | 0.6968             | 0.55655 | 0.68965             | 0.55675    | 0.6992             |
| Leaf-Size        | 0.5565625 | 0.6869375          | 0.556625 | 0.6909375          | 0.5585  | 0.6938125           | 0.5585 | 0.7005625            |

##### Concat

(1),(2) 數值一樣
(5),(6) 數值一樣

| Algorithm        | (1)       | (2)=(1).concat | (3)       | (4)=(3).concat | (5)     | (6)=(5).concat | (7)       | (8)=(7).concat |
| ---------------- | --------- | -------------- | --------- | -------------- | ------- | -------------- | --------- | -------------- |
| Pre-Order        | 0.54575   | 0.5458         | 0.69665   | 0.69985        | 0.5575  | 0.5576         | 0.68905   | 0.6936         |
| Invert-Pre-Order | 0.54435   | 0.5444         | 0.69105   | 0.6968         | 0.55655 | 0.55675        | 0.68965   | 0.6992         |
| Leaf-Size        | 0.5565625 | 0.556625       | 0.6869375 | 0.6909375      | 0.5585  | 0.5585         | 0.6938125 | 0.7005625        |

:::spoiler (1) result-cpu-sort-false-forecast-0-concat-false
```yaml=
datasetInfo:
  name: DataSet-11         # name of the dataset
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
      lambda: 900          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.54575     | 176.2    | 510.6    | 267.2    | 46       |
| Invert-Pre-Order | 0.54435     | 176.2    | 514      | 266      | 43.8     |
| Leaf-Size        | 0.5565625   | 153.5    | 519.5    | 274.25   | 52.75    |

![](https://i.imgur.com/SEdGyy5.png)

:::spoiler (2) result-cpu-sort-false-forecast-0-concat-true
```yaml=
datasetInfo:
  name: DataSet-11         # name of the dataset
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
      lambda: 900          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.5458      | 176.2    | 510.4    | 267.4    | 46       |
| Invert-Pre-Order | 0.5444      | 176.2    | 513.8    | 266.2    | 43.8     |
| Leaf-Size        | 0.556625    | 153.5    | 519.25   | 274.5    | 52.75    |

![](https://i.imgur.com/3sQefpu.png)

:::spoiler (3) result-cpu-sort-false-forecast-1-concat-false
```yaml=
datasetInfo:
  name: DataSet-11         # name of the dataset
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
      lambda: 900          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.69665     | 56.2     | 319.4    | 406      | 218.4    |
| Invert-Pre-Order | 0.69105     | 56.2     | 325.8    | 415.6    | 202.4    |
| Leaf-Size        | 0.6869375   | 59.75    | 336      | 401      | 203.25   |

![](https://i.imgur.com/m3GV7Zi.png)

:::spoiler (4) result-cpu-sort-false-forecast-1-concat-true
```yaml=
datasetInfo:
  name: DataSet-11         # name of the dataset
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
      lambda: 900          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.69985     | 54.2     | 313.8    | 410.4    | 221.6    |
| Invert-Pre-Order | 0.6968      | 54.2     | 315      | 420.2    | 210.6    |
| Leaf-Size        | 0.6909375   | 58.25    | 327.5    | 406.5    | 207.75   |

![](https://i.imgur.com/IypZHaO.png)

:::spoiler (5) result-cpu-sort-true-forecast-0-concat-false
```yaml=
datasetInfo:
  name: DataSet-11         # name of the dataset
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
      lambda: 900          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.5575      | 142.8    | 527.4    | 286.8    | 43       |
| Invert-Pre-Order | 0.55655     | 142.8    | 528.6    | 288.2    | 40.4     |
| Leaf-Size        | 0.5585      | 149.25   | 514.5    | 289.25   | 47       |

![](https://i.imgur.com/jiuU5WM.png)

:::spoiler (6) result-cpu-sort-true-forecast-0-concat-true
```yaml=
datasetInfo:
  name: DataSet-11         # name of the dataset
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
      lambda: 900          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.5576      | 142.8    | 527.4    | 286.4    | 43.4     |
| Invert-Pre-Order | 0.55675     | 142.8    | 528.6    | 287.4    | 41.2     |
| Leaf-Size        | 0.5585      | 149.25   | 514.5    | 289.25   | 47       |

![](https://i.imgur.com/3LutM5f.png)

:::spoiler (7) result-cpu-sort-true-forecast-1-concat-false
```yaml=
datasetInfo:
  name: DataSet-11         # name of the dataset
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
      lambda: 900          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.68905     | 49.8     | 331.2    | 432      | 187      |
| Invert-Pre-Order | 0.68965     | 49.8     | 329.4    | 433.2    | 187.6    |
| Leaf-Size        | 0.6938125   | 52       | 325.25   | 418.25   | 204.5    |

![](https://i.imgur.com/S33HB9c.png)

:::spoiler (8) result-cpu-sort-true-forecast-1-concat-true
```yaml=
datasetInfo:
  name: DataSet-11         # name of the dataset
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
      lambda: 900          # generate slice reqeusted cpu by poisson distribution
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
| Pre-Order        | 0.6936      | 47.4     | 322.8    | 437.8    | 192      |
| Invert-Pre-Order | 0.6992      | 47.4     | 309.4    | 442.2    | 201      |
| Leaf-Size        | 0.7005625   | 49.25    | 311.25   | 427.5    | 212      |

![](https://i.imgur.com/guWt1C0.png)