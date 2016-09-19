Pizzeria
=========

Pizzeria is a chart generator as a service


### Requirements

Golang 1.3+

### What is it?

The main inspiration of Pizzeria was
[Google deprecated chart api](https://developers.google.com/chart/image). 
We believe that sometimes developers need to have a simple way to generate
image charts. 
A http api could be a good tool to solve this problem.
Thank you for [wcharczuk/go-chart](https://github.com/wcharczuk/go-chart).


### How to install and run it ?
Prefer always a stable release [found here](https://github.com/cloudson/pizzeria/releases)

```bash
download the project to some $path
cd $path/
go run main.go
```

### Examples 

#### Pie chart

http://localhost:8080/pie?h=200&w=200&dt=1,2,3&lb=cash,credit,debit

![](./_images/pie.png)

#### Bar chart 

http://localhost:8080/bar?h=200&w=200&dt=1,2,3&lb=cash,credit,debit

![](./_images/bar.png)

### Line chart 

http://localhost:8080/line?h=200&w=200&dtx=1,2,3&dty=2,4,6

![](./_images/line.png)

### Documentation 

Read more about the possibilities on [Doc page](./doc/index.md)


