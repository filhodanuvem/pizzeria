Line
========

![](./_images/line.png)

To generate a line bar you may define the values of cartesian coordinates (x,y)

### Data X (*dtx* - required)
See below a request to generate a line chart with width=800px and hegiht=600px using 
1, 2, 3 as data of x coordinate.

`GET /line?w=800&h=600&dtx=1,2,3` 

### Data Y (*dty* - required)
It is necessary to set the values of Y coordinates

`GET /line?w=800&h=600&dtx=1,2,3&dty=2,4,6` 

### Labels (*lb* - optional)
Sometimes we want to switch the values that appear on Y axis, you can
use the label parameter to set it.

`GET /line?w=800&h=600&dtx=1,2,3&dty=2,4,6&lb=2,4,six` 

### colors (*cl* - optional)
If you want define the colors of the chart parts you can use the cl (colors) parameter.

It's important to know that you may pass the same number of colors than data. Each color
may to be on hexadecimal format without '#' prefix.

`GET /bar?w=800&h=600&dt=10,20,70&lb=cash,debit,credit&cl=f00,00ff00,00f` 
