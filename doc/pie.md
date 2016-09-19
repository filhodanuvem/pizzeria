Pie
========

![](../_images/pie.png)

### Data (*dt* - required)
To set data of the chart use the *dt* parameter using a comma char as a separator.
See below a request to generate a chart with width=800px and hegiht=600px using 
10, 20, 70 as parts of the pie.

`GET /pie?w=800&h=600&dt=10,20,70` 

### Labels (*lb* - required)
If you defined a pie chart with three type of data, it would be important named these
using the parameters lb that represents labels. 

`GET /pie?w=800&h=600&dt=10,20,70&lb=cash,debit,credit` 

### colors (*cl* - optional)
If you want define the colors of the chart parts you can use the cl (colors) parameter.

It's important to know that you may pass the same number of colors than data. Each color
may to be on hexadecimal format without '#' prefix.

`GET /pie?w=800&h=600&dt=10,20,70&lb=cash,debit,credit&cl=f00,00ff00,00f` 
