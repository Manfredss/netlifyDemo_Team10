<template>
  <div>
    <h1 class="green">Welcome!</h1>
    <p class="green">This is a test web page</p>
    <!-- <div> -->
        <img id="pic" width="840" height="600" :src="imgs">
    <!-- </div> -->

    <div>
      <button type="primary" style="font-size: 25px;" @click="PreviousImage()">Previous Image</button>
      <button type="primary" style="font-size: 25px;" @click="NextImage()">Next Image</button>
    </div>
    
  </div>
</template>

<script>
export default{
  data(){
    return{
      imgs: "assets/image1.jpg",
      imgArr:[
        "assets/image1.jpg",
        "assets/image2.jpg",
        "assets/logo-Stanford_University-2020_08_11_14_58_29.png",
        "assets/Ferrari F1-75.jpg",
        "assets/image5.jpg",
      ],
    };
  },
  methods: {
    i: 0,
    PreviousImage(){
      // var msg = document.getElementById("pic");
      if (i - 1 < 0)
        i = 4;
      img = imgArr[data().i];
      // msg.src = data.imgArr[i];
    },
    NextImage(){
      // var msg = document.getElementById("pic");
      if (i + 1 > 4)
        i = 0;
      imgs = imgArr[i];
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