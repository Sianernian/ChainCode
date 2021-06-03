package main

import (
	"fmt"
	"github.com/fabric-chaincode-go-main/shim"
	"github.com/fabric-protos-go-main/peer"
)

type HelloChainCode struct{

}

func (h *HelloChainCode)Init(stub shim.ChaincodeStubInterface)peer.Response{
	fmt.Println("开始实例化链码")
	_,args:=stub.GetFunctionAndParameters()
	if len(args) !=2{
		return shim.Error("参数错误")
	}
	fmt.Println("保存数据中")
	return shim.Success(nil)

}

func (h *HelloChainCode)Invoke(stub shim.ChaincodeStubInterface)peer.Response{
	fun,args:=stub.GetFunctionAndParameters()
	if fun =="query"{
		return query(stub,args)
	}
	return shim.Error("操作违法，无法执行")

}

func query(stub shim.ChaincodeStubInterface , args []string)peer.Response{
	if len(args)!=1{
		return shim.Error("参数有错，必须为指定参数")
	}
	result,err :=stub.GetState(args[0])
	if err !=nil{
		return shim.Error(args[0]+"查询是错误")
	}
	if result ==nil{
		return shim.Error(args[0]+"没有数据")
	}
	return shim.Success(result)
}


func main() {
	err :=shim.Start(new(HelloChainCode))
	if err!=nil{
		fmt.Println("启动HelloChainCode失败，请重试:%s",err)
	}

}



