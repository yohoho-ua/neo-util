new Vue ({
    el: '#app',
    mounted: function () {
        this.fetchAssets();
    },
    data: {
    birds: [],
    address: '',
    currency: ['NEO', 'NeoGas'],
    sum: 0,
    neoAsset: "",
    gasAsset: "",
    transactionMade: false,
    transactionStatus: true,
    transactionMessage: '',
    statusClass: '',
    checkAddressField: 'form-control',
    checkValueField: 'form-control',
    checkSumField: 'form-control'
    },
    methods:{
    //     transfer: function(){
    //         if (this.address.length!=34){
    //         this.checkAddressField = 'form-control is-invalid'}
    //         else if (this.sum<=0){
    //             this.checkSumField = 'form-control is-invalid'
    //         }
    //         else if (this.currency!='NEO'&&this.currency!='NeoGas'){
    //             this.checkValueField = 'form-control is-invalid'
    //         }
    //         else{
    //  //  if (this.address.length==34 && this.sum>0 && (this.currency=='NEO'||this.currency=='NeoGas')) {
    //     this.transactionMade=true;
    //     this.checkAddressField = 'form-control',
    //     this.checkValueField = 'form-control',
    //     this.checkSumField = 'form-control'
    //     this.checkField = 'form-control'
    //     if (this.transactionStatus == true){
    //         this.statusClass = 'text-success lead'
    //         this.transactionMessage = 'Transaction maid!'
    //     }
    //     else{
    //         this.statusClass = 'text-danger lead'
    //         this.transactionMessage = 'Transaction failed!'
    //     }
    //    }}
            fetchAssets: function () {
                var transactions = [];
                fetch("/neo")
                .then(response => response.json())
                .then(json => {
                console.log(json.neo, json.gas);
                this.neoAsset = json.neo;
                this.gasAsset = json.gas;
                })
              },

        },


    }
);