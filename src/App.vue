<template>
  <div>
    <p style="font-size: 39px; left: -10px; top: -120px; font-weight:500">COMP2002 Team 10 Use of Color: Prototype Demo</p>
    <img id="pic" width="840" height="600" :src="imgs">

    <div>
      <button type="primary" style="font-size: 25px;" @click="PreviousImage">Previous Image</button>
      <button type="primary" style="font-size: 25px; left: 507px;" @click="NextImage">Next Image</button>
    </div>
  </div>

  <div>
      <p>Q1. What's the image type?<br>
        <input type="radio" id="1" name="Q1" style="left: 60px"><label for="1" style="left:70px;">Black & Whtie</label>
        <input type="radio" id="2" name="Q1" style="left:100px;"><label for="2" style="left:110px;">Greyscale</label>
        <input type="radio" id="3" name="Q1" style="left:140px;"><label for="3" style="left:150px;">Colour</label>
      </p>
      <br>

      <p>Q2. Is the color used for Aesthetics or Data Visualisation?<br>
        <input type="radio" id="1" name="Q2" style="left: 60px"><label for="1" style="left:70px;">Aesthetics</label>
        <input type="radio" id="2" name="Q2" style="left:100px;"><label for="2" style="left:110px;">Data Visualisation</label>
        <input type="radio" id="3" name="Q2" style="left:140px;"><label for="3" style="left:150px;">Not sure</label>
      </p>
      <br>

      <p>Q3. Is a colour mapping legend showed?<br>
        <input type="radio" id="1" name="Q3" style="left: 60px"><label for="1" style="left:70px;">Yes</label>
        <input type="radio" id="2" name="Q3" style="left:100px;"><label for="2" style="left:110px;">No</label>
        <input type="radio" id="3" name="Q3" style="left:140px;"><label for="3" style="left:150px;">Not sure</label>
      </p>
      <br>

      <p>Q4. How many colours are used in this image?<br>
        <textarea style="font-size: 25px; text-align: center; resize: none" placeholder="Enter the number" :rows="1"></textarea>
        <input type="radio" id="1" style="left:140px;"><label for="1" style="left:150px;">Not sure</label>
      </p>
      <br>

      <p>Q5. Which colour mapping is used, continuous or categorical?<br>
        <input type="radio" id="1" name="Q5" style="left: 60px"><label for="1" style="left:70px;">continuous</label>
        <input type="radio" id="2" name="Q5" style="left:100px;"><label for="2" style="left:110px;">Categorical</label>
        <input type="radio" id="3" name="Q5" style="left:140px;"><label for="3" style="left:150px;">Not sure</label>
      </p>
      <br>

      <p>Q6. From 1-5, how would you rank the difficulty of identifying this image?<br>
        <input type="radio" id="1" name="Q6" style="left: 60px"><label for="1" style="left:70px;">1</label>
        <input type="radio" id="2" name="Q6" style="left:100px;"><label for="2" style="left:110px;">2</label>
        <input type="radio" id="3" name="Q6" style="left:140px;"><label for="3" style="left:150px;">3</label>
        <input type="radio" id="4" name="Q6" style="left:180px;"><label for="4" style="left:190px;">4</label>
        <input type="radio" id="5" name="Q6" style="left:220px;"><label for="5" style="left:230px;">5</label>
      </p>
      <button style="font-size: 30px; left: 400px; top: 10px;">submit</button>
  </div>
</template>

<script>
import urlJson from "/url.json";
export default{
  data(){
    return{
      i : 0,
      imgs: "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.6.1.png",
      imgArr: urlJson,
    };
  },

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
    top: -60px;
}p{
  font-weight: 450;
  font-size: 1.5rem;
  top: -10px;
  left: 60px;
}label{
  position: relative
}textarea{
  width: 240px;
  height: 50px;
  left: 90px; 
  top: 15px;
  border-radius: 6px;
  font: Arial;
}button:hover{
  border-radius: 5px;
  -webkit-box-shadow: 3px 3px 3px rgba(0,0,0, .29), inset 1px 1px 1px rgba(255,255,255, .44);
  
  -webkit-transition: all 0.15s ease;  
  transition: all 0.15s ease; 
}
input[type="radio"]{
  width: 15px;
  height: 15px;
}
</style>

<!-- -moz-box-shadow: 3px 3px 3px rgba(0,0,0, .29), inset 1px 1px 1px rgba(255,255,255, .44);   -->
  <!-- box-shadow: 3px 3px 3px rgba(0,0,0, .29), inset 1px 1px 1px rgba(255,255,255, .44);   -->
  <!-- -moz-transition: all 0.15s ease; 
  -o-transition: all 0.15s ease;  
  -ms-transition: all 0.15s ease; -->

<!-- fetch('./url.json')
          .then(function(response) {
            return response.json().then(function(json){
              for(i = 0; i < json.length(); i++){
                var url = json[i].url;
                this.imgArr.push(url);
              }
            });
          }) -->
