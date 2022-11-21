<template>
  <div>
    <h1 class="green">Welcome!</h1>
    <p class="green">This is a test web page</p>
    <!-- <div> -->
        <img id="pic" width="840" height="600" :src="imgs">
    <!-- </div> -->

    <div>
      <button type="primary" style="font-size: 25px;" @click="PreviousImage">Previous Image</button>
      <button type="primary" style="font-size: 25px;" @click="NextImage">Next Image</button>
    </div>
    
  </div>
</template>

<script>
export default{
  data(){
    return{
      i : 0,
      imgs: "/src/assets/image1.jpg",
      imgArr:[
        "/src/assets/image1.jpg",
        "/src/assets/image2.jpg",
        "/src/assets/logo-Stanford_University-2020_08_11_14_58_29.png",
        "/src/assets/Ferrari F1-75.jpg",
        "/src/assets/image5.jpg",
      ],
    };
  },
  methods: {
    PreviousImage(){
      // var msg = document.getElementById("pic");
      if (this.i - 1 < 0)
        this.i = 4;
      console.log(this.i);
      imgs = imgArr[this.i];
      // msg.src = data.imgArr[i];
    },
    NextImage(){
      // var msg = document.getElementById("pic");
      if (this.i + 1 > 4)
        this.i = 0;
      console.log(this.i);
      imgs = imgArr[this.i];
      // msg.src = data.imgArr[i];
    }
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
