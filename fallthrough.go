package main

import "fmt"

func main() {

	switch {
	case false: //此时case值空为0即false，所以不执行1
		fmt.Println("1、case 条件语句为 false")
		fallthrough //使用 fallthrough 会强制执行后面的 case 语句
	case true: //1非空即为真，执行2，有fall through继续判断
		fmt.Println("2、case 条件语句为 true")
		fallthrough //fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	case false: //真，执行3，有fall through继续判断
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true: //真，执行4，无fall through不再继续判断5
		fmt.Println("4、case 条件语句为 true")
	case false: //前面满足case条件，所以5不执行
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}

/*
switch{
    case 1:
    ...
    if(...){
        break
    }

    fallthrough // 此时switch(1)会执行case1和case2，但是如果满足if条件，则只执行case1

    case 2:
    ...
    case 3:
}


switch 的 default 不论放在哪都是最后执行：

=====================================================================
a := 10

switch {
   default : {
      fmt.Println("default")
   }
   case a > 0 : {
      fmt.Println("a > 0")
   }
   case a >5 : {
      fmt.Println("a > 5")
   }
}

=======================================================================
a := 10
switch {
   case a > 0 : {
      fmt.Println("a > 0")
   }
   case a >5 : {
      fmt.Println("a > 5")
   }
               default : {
      fmt.Println("default")
   }
}

*/
