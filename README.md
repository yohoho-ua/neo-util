# neo-util created for research and study purpose
Test app for sending NEO and GAS    
![alt text](https://i.imgur.com/ibJDpPR.png "Kartinochka")


Running rpc client with open wallet required. Use neo-cli
```dotnet neo-cli.dll /rpc
   open wallet <path_to_wallet>
```  
* enter your pass

* edit conf.json for account address, and NEO-node rpc host  

```
const (
	accountAddress = "your_address_here"
)
```

by default app is running on [http://localhost:8080/assets/](http://localhost:8080/assets/ "link to wallet")