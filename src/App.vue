<template>
  <div>
    <h1 class="green">Welcome!</h1>
    <p class="green">This is a demo web page</p>
    <img id="pic" width="840" height="600" :src="imgs">

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
      imgs: "http://dx.doi.org/10.1109/VISUAL.1990.146359",
      imgArr:[
        "http://dx.doi.org/10.1109/VISUAL.1990.146359",
        "http://dx.doi.org/10.1109/VISUAL.1990.146359",
        "/src/assets/logo-Stanford_University-2020_08_11_14_58_29.png",
        "http://dx.doi.org/10.1109/VISUAL.1990.146359",
        "http://dx.doi.org/10.1109/VISUAL.1990.146359",
      ],
    };
  },
  methods: {
    PreviousImage: function(){
      if (this.i < 0)
        this.i = 4;
      this.imgs = this.imgArr[this.i];
      this.i -= 1;
    },
    NextImage: function(){
      if (this.i > 4)
        this.i = 0;
      this.imgs = this.imgArr[this.i];
      this.i += 1;
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
