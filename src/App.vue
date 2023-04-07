<template>
  <div id="backgroundimg">
    <!-- <div id="user">
        <label style="font-size: 36px; left: 740px; top: -90px;">User name:</label>
        <select type="text" style="font-size: 36px; left: 750px; top: -90px; text-align: center;" placeholder="Select User" list="typelist" v-model="username">
            <option>Robert Laramee</option>
            <option>Jamie Vickers</option>
            <option>Zhening Zhu</option>
            <option>Luke Whitfield</option>
            <option>Wenfei Qi</option>
            <option>Nita Krasniqi</option>
            <option>Yizhan Huang</option>
            <option>Hyanggi Lee</option>
        </select>
        <button type="primary" style="font-size: 36px; top: -20px; left: 496px;" @click="Login">Log in</button>
      </div> -->

    <div id="header">
      <p style="font-size: 60px; left: 30px; top: 10px; font-weight:500">CO10UR Use of Color</p>
      <n-space style="left:700px; top:-51px;">
        <n-button strong round size="large" type="primary" style="font-size:45px;" @click="Annotation" ghost>Annotation</n-button>
        <n-button strong round size="large" type="primary" style="font-size:45px;" @click="Validation" ghost>Validation</n-button>
        <n-button strong round size="large" type="primary" style="font-size:45px;" @click="Exploration" ghost>Exploration</n-button>
      </n-space>
      <n-button type="info" size="large" style="left:1900px; top:-99px; font-size:45px; font-weight:bold;" dashed>Log Out</n-button>
      <p style="top: -90px; left: 30px; font-size: 51px; color: coral">Welcome! {{username}}</p>
    </div>

    <div id="annotation" style="top:-300px;"> <!-- style="display:none">-->
      <n-grid x-gap="200px" :cols="2">
        <n-gi style="left: 70px; top: 300px;">
          <n-card id="imgcard" size="huge">
            <div style="top:-40px; left:21px;">
              <div id="imgg" style="top:30px;">
                <a :href="imgs" style="font-size: 21px;">{{imgs}}</a>
                <img id="pic" width="840" height="600" :src="imgs">
              </div>

              <div style="top:45px;">
                <button type="primary" style="font-size: 25px;" @click="PreviousImage">Previous Image</button>
                <input type="text" :src="pageNumber" v-model="pageNumber" 
                       style="left:120px; width:120px; font-size:25px; height:36px; 
                       text-align: center; resize: none" :rows="1" @keyup.enter="dokeyup">
                <label style="left:130px; font-size:25px;">/ 29689</label>
                <button type="primary" style="font-size: 25px; left: 303px;" @click="NextImage">Next Image</button>
              </div>
            </div>
          </n-card>
        </n-gi>
        
        <n-gi style="left:-100px; top:300px;">
          <div id="choose"> <!--style="display:none; top:110px; left:60px;">-->
            <n-space vertical>
              <n-card>
                <label style="font-size:24px; left:30px;">Q1. What's the image type?<br></label>
                  <input type="radio" id="1" v-model="Q1" style="left: 60px;" value="Black & White"><label for="1" style="left:70px;">Black & White</label>
                  <input type="radio" id="2" v-model="Q1" style="left:100px;" value="Greyscale"><label for="2" style="left:110px;">Greyscale</label>
                  <input type="radio" id="3" v-model="Q1" style="left:140px;" value="Colour"><label for="3" style="left:150px;">Colour</label>
              </n-card>

              <n-card>
                <label style="font-size:24px; left:30px;">Q2. Is the color used for Aesthetics or Data Visualisation?<br></label>
                  <n-radio-group style="font-size:21px;">
                    <input type="radio" id="1" v-model="Q2" style="left: 60px;" value="Aesthetics"><label for="1" style="left:70px;">Aesthetics</label>
                    <input type="radio" id="2" v-model="Q2" style="left:100px;" value="Color Mapping"><label for="2" style="left:110px;">Color Mapping</label>
                    <input type="radio" id="2" v-model="Q2" style="left:140px;" value="Depth Perception"><label for="3" style="left:150px;">Depth Perception</label>
                    <input type="radio" id="3" v-model="Q2" style="left:180px;" value="Not Applicable"><label for="4" style="left:190px;">Not Applicable</label>
                    <input type="radio" id="4" v-model="Q2" style="left:220px;" value="Not Sure"><label for="5" style="left:230px;">Not Sure</label>
                  </n-radio-group>
              </n-card>
              
              <n-card>
                <label style="font-size:24px; left:30px;">Q3. Is a colour mapping legend showed?<br></label>
                    <n-button-group>
                      <input type="radio" id="1" v-model="Q3" style="left: 60px;" value="Yes"><label for="1" style="left:70px;">Yes</label>
                      <input type="radio" id="2" v-model="Q3" style="left:100px;" value="No"><label for="2" style="left:110px;">No</label>
                      <input type="radio" id="3" v-model="Q3" style="left:140px;" value="Not Sure"><label for="3" style="left:150px;">Not sure</label>
                    </n-button-group>
              </n-card>

              <n-card>
                <label style="font-size:24px; left:30px;">Q4. How many colours are used in this image?<br></label>
                  <input type="textarea" id="1" v-model="Q4" style="font-size: 25px; text-align: center; resize: none" placeholder="Enter the number" :rows="1">
                  <input type="radio" id="2" v-model="Q4" style="left:75px;" value="Not Sure"><label for="1" style="left:85px;">Not sure</label>
                  <input type="radio" id="3" v-model="Q4" style="left:140px;" value="Not Applicable"><label for="2" style="left:150px;">Not applicable</label>
              </n-card>

              <n-card>
                <label style="font-size:24px; left:30px;">Q5. Which colour mapping is used, continuous or categorical?<br></label>
                  <input type="radio" id="1" v-model="Q5" style="left: 60px;" value="Continuous"><label for="1" style="left:70px;">Continuous</label>
                  <input type="radio" id="2" v-model="Q5" style="left:100px;" value="Categorical"><label for="2" style="left:110px;">Categorical</label>
                  <input type="radio" id="3" v-model="Q5" style="left:140px;" value="Both"><label for="3" style="left:150px;">Both</label>
                  <input type="radio" id="4" v-model="Q5" style="left:180px;" value="Not Applicable"><label for="4" style="left:190px;">Not applicable</label>
              </n-card>

              <n-card>
                <label style="font-size:24px; left:30px;">Q6. From 1-5 (easy-difficult), how would you rank the difficulty of identifying this image?<br></label>
                  <n-radio-group v-model:value="Q6"  size="large">
                    <n-radio value="1" style="font-size:22px; left:100px;">1</n-radio>
                    <n-radio value="2" style="font-size:22px; left:150px;">2</n-radio>
                    <n-radio value="3" style="font-size:22px; left:200px;">3</n-radio>
                    <n-radio value="4" style="font-size:22px; left:250px;">4</n-radio>
                    <n-radio value="5" style="font-size:22px; left:300px;">5</n-radio>
                  </n-radio-group>
              </n-card>
              <button style="font-size: 30px; left: 400px; top: 10px;" @click="Submit">Annotate</button>
            </n-space>
          </div>
        </n-gi>
      </n-grid>
    </div>
  </div>
</template>

<script>
import urlJson from "/url.json";
import axios from 'axios';
import qs from 'qs';
export default{
  data(){
    return{
      i : 0,
      imgs: "https://web.cse.ohio-state.edu/~chen.8028/VisPubImages/Images/1990/VisC.6.1.png",
      imgArr: urlJson,
      Q1:'',
      Q2:'',
      Q3:'',
      Q4:'',
      Q5:'',
      Q6:'',
      pageNumber: 1,
      username:'',
    };
  },

  methods: {
    Login: function(){
      (document.getElementById("user")).style.display = "none";
      (document.getElementById("annotation")).style.display = "block";
      // (document.getElementById("choose")).style.display = "block";
    },
    Annotation: function(){
      (document.getElementById("annotation")).style.display = "block";
      // (document.getElementById("choose")).style.display = "block";
    },
    Validation: function(){
      (document.getElementById("annotation")).style.display = "none";
      // (document.getElementById("choose")).style.display = "none";
    },
    Exploration: function(){
      (document.getElementById("annotation")).style.display = "none";
      // (document.getElementById("choose")).style.display = "none";
    },
    PreviousImage: function(){
      this.i -= 1;
      if (this.i < 0)
        this.i = this.imgArr.length - 1;
      this.pageNumber = this.imgArr[this.i].imageID;
      this.imgs = this.imgArr[this.i].url;

      axios.get('http://127.0.0.1:8080/get', {
        params:{imageID: this.imgArr[this.i].imageID}
      }).then((res) =>{
        if(res.status === 200){
          this.Q1 = res.data.Q1;
          this.Q2 = res.data.Q2;
          this.Q3 = res.data.Q3;
          this.Q4 = res.data.Q4;
          this.Q5 = res.data.Q5;
          this.Q6 = res.data.Q6;
        }
        else{
          alert(res.data)
        }
      })
    },
    NextImage: function(){
      const data = {
        userName: this.username,
        imageID: this.imgArr[this.i].imageID,
        Q1: this.Q1,
        Q2: this.Q2,
        Q3: this.Q3,
        Q4: this.Q4,
        Q5: this.Q5,
        Q6: this.Q6,
      };

      axios.post('http://127.0.0.1:8080/save', data)
      .then((res) => {
        if(res.status === 200){
          alert("Record saved successfully!")
        }
        else {
          alert(res.data)
        }
      });

      this.i += 1;
      if (this.i > this.imgArr.length - 1)
        this.i = 0;
      this.pageNumber = this.imgArr[this.i].imageID;
      this.imgs = this.imgArr[this.i].url;

      axios.get('http://127.0.0.1:8080/get',{
        params:{imageID: this.imgArr[this.i].imageID}
      }).then((res) =>{
        if(res.status === 200){
          this.Q1 = res.data.Q1;
          this.Q2 = res.data.Q2;
          this.Q3 = res.data.Q3;
          this.Q4 = res.data.Q4;
          this.Q5 = res.data.Q5;
          this.Q6 = res.data.Q6;
        }
        else{
          alert(res.data)
        }
      })
    },
    Submit: function(){
      const data = {
        userName: this.username,
        imageID: this.imgArr[this.i].imageID,
        Q1: this.Q1,
        Q2: this.Q2,
        Q3: this.Q3,
        Q4: this.Q4,
        Q5: this.Q5,
        Q6: this.Q6,
      };

      axios.post('http://127.0.0.1:8080/save', data)
      .then((res) => {
        if(res.status === 200){
          alert("Record saved successfully!")
        }
        else {
          alert(res.data)
        }
      });
    },
    dokeyup: function(e){
      if(this.pageNumber === "" || this.pageNumber == null)
          alert("Number can't be empty!");
      else if(isNaN(this.pageNumber))
          alert("Invalid input! Please enter a number!");
      else if(this.pageNumber < 1 || this.pageNumber > 29689)
          alert("Index out of range! Number should between 1 and 29689 (1 and 29689 included).");
      else{
          this.imgs = this.imgArr[this.pageNumber - 1].url;
          this.i = this.pageNumber - 1;
      }

      axios.get('http://127.0.0.1:8080/get',{
        params:{imageID: this.imgArr[this.i].imageID}
      }).then((res) =>{
        if(res.status === 200){
          this.Q1 = res.data.Q1;
          this.Q2 = res.data.Q2;
          this.Q3 = res.data.Q3;
          this.Q4 = res.data.Q4;
          this.Q5 = res.data.Q5;
          this.Q6 = res.data.Q6;
        }
        else{
          alert(res.data)
        }
      })
    },
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
  
  -webkit-transition: all 0.15s ease;  
  transition: all 0.15s ease; 
}
input[type="radio"]{
  width: 15px;
  height: 15px;
}
#backgroundimg{
  background: url("assets/background_img.jpg") no-repeat;
  background-position: center;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  background-size: cover;
  position: fixed;
}
</style>