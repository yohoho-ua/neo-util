# neo-util
Test app for sending NEO and GAS    
![alt text](https://i.imgur.com/8kIFvVp.png "Kartinochka")


Running rpc client with open wallet required. Use neo-cli
```dotnet neo-cli.dll /rpc
   open wallet <path_to_waallet>
```  
* enter your pass

* edit rpcClient.go, line: 11  

```
const (
	accountAddress = "your_address_here"
)
```

