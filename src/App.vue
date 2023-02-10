<template>
  <div>
    <p style="font-size: 39px; left: -10px; top: -120px; font-weight:500">COMP2002 Team 10 Use of Color: Prototype Demo</p>
    <a :href="imgs" style="font-size: 21px;">{{imgs}}</a>
    <img id="pic" width="840" height="600" :src="imgs">
    <div>
      <button type="primary" style="font-size: 25px;" @click="PreviousImage">Previous Image</button>
      <input type="text" :src="pageNumber" v-model="pageNumber" style="left:120px; width:120px; font-size:25px; height:36px; text-align: center; resize: none" :rows="1" @keyup.enter="dokeyup">
      <label style="left:130px; font-size:25px;">/ 29689</label>
      <button type="primary" style="font-size: 25px; left: 303px;" @click="NextImage">Next Image</button>
    </div>
  </div>

  <div id="choose">
      <p>Q1. What's the image type?<br>
        <input type="radio" id="1" v-model="Q1" style="left: 60px;" value="Black & White"><label for="1" style="left:70px;">Black & White</label>
        <input type="radio" id="2" v-model="Q1" style="left:100px;" value="Greyscale"><label for="2" style="left:110px;">Greyscale</label>
        <input type="radio" id="3" v-model="Q1" style="left:140px;" value="Colour"><label for="3" style="left:150px;">Colour</label>
      </p>
      <br>

      <p>Q2. Is the color used for Aesthetics or Data Visualisation?<br>
        <input type="radio" id="1" v-model="Q2" style="left: 60px;" value="Aesthetics"><label for="1" style="left:70px;">Aesthetics</label>
        <input type="radio" id="2" v-model="Q2" style="left:100px;" value="Data Visualisation"><label for="2" style="left:110px;">Data Visualisation</label>
        <input type="radio" id="3" v-model="Q2" style="left:140px;" value="Not Sure"><label for="3" style="left:150px;">Not sure</label>
      </p>
      <br>

      <p>Q3. Is a colour mapping legend showed?<br>
        <input type="radio" id="1" v-model="Q3" style="left: 60px;" value="Yes"><label for="1" style="left:70px;">Yes</label>
        <input type="radio" id="2" v-model="Q3" style="left:100px;" value="No"><label for="2" style="left:110px;">No</label>
        <input type="radio" id="3" v-model="Q3" style="left:140px;" value="Not Sure"><label for="3" style="left:150px;">Not sure</label>
      </p>
      <br>

      <p>Q4. How many colours are used in this image?<br>
        <input type="textarea" id="1" v-model="Q4" style="font-size: 25px; text-align: center; resize: none" placeholder="Enter the number" :rows="1">
        <input type="radio" id="2" v-model="Q4" style="left:140px;" value="Not Sure"><label for="1" style="left:150px;">Not sure</label>
      </p>
      <br>

      <p>Q5. Which colour mapping is used, continuous or categorical?<br>
        <input type="radio" id="1" v-model="Q5" style="left: 60px;" value="Continuous"><label for="1" style="left:70px;">Continuous</label>
        <input type="radio" id="2" v-model="Q5" style="left:100px;" value="Categorical"><label for="2" style="left:110px;">Categorical</label>
        <input type="radio" id="3" v-model="Q5" style="left:140px;" value="Not Sure"><label for="3" style="left:150px;">Not sure</label>
      </p>
      <br>

      <p>Q6. From 1-5, how would you rank the difficulty of identifying this image?<br>
        <input type="radio" id="1" v-model="Q6" style="left: 60px;" :value="1"><label for="1" style="left:70px;">1</label>
        <input type="radio" id="2" v-model="Q6" style="left:100px;" :value="2"><label for="2" style="left:110px;">2</label>
        <input type="radio" id="3" v-model="Q6" style="left:140px;" :value="3"><label for="3" style="left:150px;">3</label>
        <input type="radio" id="4" v-model="Q6" style="left:180px;" :value="4"><label for="4" style="left:190px;">4</label>
        <input type="radio" id="5" v-model="Q6" style="left:220px;" :value="5"><label for="5" style="left:230px;">5</label>
      </p>
      <button style="font-size: 30px; left: 400px; top: 10px;" @click="Submit">submit</button>
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
      Q1:'',
      Q2:'',
      Q3:'',
      Q4:'',
      Q5:'',
      Q6:'',
      pageNumber: 1,
    };
  },

  methods: {
    PreviousImage: function(){
      this.i -= 1;
      if (this.i < 0)
        this.i = this.imgArr.length - 1;
      this.pageNumber = this.imgArr[this.i].imageID;
      this.imgs = this.imgArr[this.i].url;
      
    },
    NextImage: function(){
      this.i += 1;
      if (this.i > this.imgArr.length - 1)
        this.i = 0;
      this.pageNumber = this.imgArr[this.i].imageID;
      this.imgs = this.imgArr[this.i].url;
    },
    Submit: function(){
      alert(this.Q1 + " " + this.Q2 + " " + this.Q3 + " " + this.Q4 + " " + this.Q5 + " " + this.Q6);
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
  -webkit-box-shadow: 3px 3px 3px rgba(0,0,0, .29), inset 1px 1px 1px rgba(255,255,255, .44);
  
  -webkit-transition: all 0.15s ease;  
  transition: all 0.15s ease; 
}
input[type="radio"]{
  width: 15px;
  height: 15px;
}
</style>