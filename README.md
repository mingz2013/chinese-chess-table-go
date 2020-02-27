# chinese-chess-table-go
chinese chess table go

象棋桌子逻辑，不依赖其他模块，只依赖标准库










# chess

4bits

| type | color | desc | 
|:---:|:---:|:---|
| 0b000 | 0b0 | 空|
| 0b001 | 0b0 | 車|
| 0b010 | 0b0 | 马|
| 0b011 | 0b0 | 炮|
| 0b100 | 0b0 | 象|
| 0b101 | 0b0 | 士|
| 0b110 | 0b0 | 将|
| 0b111 | 0b0 | 卒|
| 0b001 | 0b1 | 車|
| 0b010 | 0b1 | 马|
| 0b011 | 0b1 | 炮|
| 0b100 | 0b1 | 象|
| 0b101 | 0b1 | 士|
| 0b110 | 0b1 | 将|
| 0b111 | 0b1 | 卒|



# xy

9 * 10

1-9
1-10

4bits

| num | desc |
|:---:|:---:|
| 0b0000 |0|
| 0b0001 |1|
| 0b0010 |2|
| 0b0011 |3|
| 0b0100 |4|
| 0b0101 |5|
| 0b0110 |6|
| 0b0111 |7|
| 0b1000 |8|
| 0b1001 |9|
| 0b1010 | 10 |




# point

8bits

| x | y |
|:---:|:---:|
| 0b0000 | 0b0000 |


# action

16bits

|sx | sy | dx| dy | 
|:---:|:---:|:---:|:---:|
|0b0000 | 0b0000 | 0b0000 | 0b0000 |




# board

4 * 9 * 10 = 360bits = 45bytes

4 * 11 * 12 = 528bits = 66bytes

|1|1|1|1|1|1|1|1|1|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 |
| 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 |
| 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 |
| 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 |
| 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 |
| 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 |
| 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 |
| 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 |
| 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 |
| 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 | 0b0000 |



```
||     |       |       |       |       |       |       |       |       |     ||
||---------------------------（帅）-----------（士）---（象）---（马）---（車）---||
||     |       |       |       |       |       |       |       |       |     ||
||     |       |       |       |       |       |       |       |       |     ||
||-----------------------------------（士）-----------------------------------||
||     |       |       |       |       |       |       |       |       |     ||
||     |       |       |       |       |       |       |       |       |     ||
||---（車）---（炮）---（马）-----------（象）-------------------（炮）-----------||
||     |       |       |       |       |       |       |       |       |     ||
||     |       |       |       |       |       |       |       |       |     ||
||---（兵）-----------（兵）-----------（兵）-----------（兵）-----------（兵）---||
||     |       |       |       |       |       |       |       |       |     ||
||     |       |       |       |       |       |       |       |       |     ||
||---------------------------------------------------------------------------||
||     |       |       |       |       |       |       |       |       |     ||
||     |       |       |       |       |       |       |       |       |     ||
||---------------------------------------------------------------------------||
||     |       |       |       |       |       |       |       |       |     ||
||     |       |       |       |       |       |       |       |       |     ||
||---【兵】-----------【兵】-----------【兵】-----------【兵】-----------【兵】---||
||     |       |       |       |       |       |       |       |       |     ||
||     |       |       |       |       |       |       |       |       |     ||
||---【马】---【炮】-------------------------------------------【炮】-----------||
||     |       |       |       |       |       |       |       |       |     ||
||     |       |       |       |       |       |       |       |       |     ||
||---------------------------------------------------------------------------||
||     |       |       |       |       |       |       |       |       |     ||
||     |       |       |       |       |       |       |       |       |     ||
||---【車】-----------【象】---【士】---【帅】---【士】---【象】---【马】---【車】---||
||     |       |       |       |       |       |       |       |       |     ||
```