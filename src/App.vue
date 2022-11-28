<template>
  <div>
    <h1 class="green">Welcome!</h1>
    <p class="green">Click the "Load  URL" button before browsing!</p>
    <img id="pic" width="840" height="600" :src="imgs">

    <div>
      <button type="primary" style="font-size: 25px;" @click="PreviousImage">Previous Image</button>
      <button type="primary" style="font-size: 25px;" @click="NextImage">Next Image</button>
      <button type="primary" style="font-size: 25px;" @click="LoadURL">Load URL</button>
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
// [
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.6.1.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.6.3.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.6.2.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.6.6.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.6.5.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.6.4.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.6.7.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.6.8.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.6.9.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.6.10.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.6.11.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.14.9.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.14.1.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.14.3.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.14.5.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.14.6.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.14.2.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.14.4.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.14.7.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.14.8.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.28.1.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.36.1.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.36.2.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.36.4.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.36.5.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.36.3.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.45.1.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.45.3.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.45.2.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.45.7.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.45.6.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.45.9.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.45.8.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.45.4.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.45.5.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.45.10.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.51.1.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.51.2.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.51.6.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.51.4.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.51.3.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.51.7.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.51.5.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.51.8.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.59.2.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.59.1.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.59.3.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.59.5.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.59.4.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.59.6.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.59.8.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.59.7.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.59.9.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.59.11.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.59.10.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.59.14.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.59.13.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.59.12.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.59.15.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.67.5.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.67.1.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.67.6.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.67.2.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.67.3.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.67.4.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.67.7.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.67.8.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.67.9.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.67.10.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.67.11.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.67.12.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.67.13.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.67.14.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.74.3.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.74.1.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.74.2.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.74.4.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.74.5.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.74.6.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.74.8.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.74.7.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.74.9.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.83.1.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.83.3.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.83.2.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.83.4.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.83.5.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.83.6.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.83.7.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.83.8.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.83.10.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.83.9.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.83.11.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.83.12.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.83.15.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.83.13.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.83.18.png",
      // "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/new/VisC.83.14.png",
      // ],
