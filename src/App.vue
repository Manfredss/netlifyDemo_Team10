<template>
  <div>
    <h1 class="green">Welcome!</h1>
    <p class="green">This is a newer version of Image browser! </p>
    <img id="pic" width="840" height="600" :src="imgs">

    <div>
      <button type="primary" style="font-size: 25px;" @click="PreviousImage">Previous Image</button>
      <button type="primary" style="font-size: 25px;" @click="NextImage">Next Image</button>
      <!-- <button type="primary" style="font-size: 25px;" @click="LoadURL">Load URL</button> -->
  </div>
  </div>
</template>

<script>
import urlJson from "/url.json"
export default{
  data(){
    return{
      i : 0,
      imgs: "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.6.1.png",
      imgArr: urlJson,
      // [{
      //   imageID: null,
      //   name: null,
      //   url: null,
      //   doi: null,
      //   year: null,
      // }],
    };
  },
  // mounted() {
  //   axios
  //       .get("./url.json")
  //       .then(response => (this.imgArr = response.data));
  // },
  methods: {
    PreviousImage: function(){
      if (this.i < 0)
        this.i = this.imgArr.length - 1;
      this.imgs = this.imgArr[this.i].url;
      this.i -= 1;
    },
    NextImage: function(){
      if (this.i > this.imgArr.length - 1)
        this.i = 0;
      this.imgs = this.imgArr[this.i].url;
      this.i += 1;
    },
    // LoadURL: function(){
      // var url = [];
      // fetch('./url.json')
      //     .then((response) => response.json())
      //     .then((json) => this.imgArr = json);
      // for(var i = 0; i < url.length(); i++)
      //     this.imgArr.concat(url[i].url);
    //   var url = "https://files.catbox.moe/7dvpgw.json"/*json文件url，本地的就写本地的位置，如果是服务器的就写服务器的路径*/
    //   var request = new XMLHttpRequest();
    //   request.open("get", url);/*设置请求方法与路径*/
    //   request.send(null);/*不发送数据到服务器*/
    //   request.onload = function () {/*XHR对象获取到返回信息后执行*/
    //     if (request.status == 200) {/*返回状态为200，即为数据获取成功*/
    //       console.log("success");
    //       var json = JSON.parse(request.responseText);
    //       for(var i = 0;i < json.length; i++)
    //         this.imgArr.push(json[i].url);
    //     }else{
    //       console.log("fail");
    //     }
    //   }
    // },
  },
}
</script>

<style scoped>
h1{
    font-weight: 500;
    font-size: 2.6rem;
    top: -10px;
}p{
  font-weight: 500;
  font-size: 1.8rem;
  top: -10px;
}
</style>

<!-- fetch('./url.json')
          .then(function(response) {
            return response.json().then(function(json){
              for(i = 0; i < json.length(); i++){
                var url = json[i].url;
                this.imgArr.push(url);
              }
            });
          }) -->

