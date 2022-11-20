<!-- <template>
  <div>
    <div class="zs-adv">
      <a title="Previous" :href="'#'" class="adv-pre" @click="leftScroll">Previous</a>
      <div id="adv-pad-scroll">
        <div class="adv-pad">
          <img
              class="adv-pad-item"
              v-for="(item, index) in itemlist"
              :key="index"
              alt=""
              :ref="`item${index}`"
              :src="item.src"
          />
        </div>
      </div>
      <a title="Next" :href="'#'" class="adv-next" @click="rightScroll">next</a>
    </div>
  </div>
</template>

<script>
export default{
  name: "index",
  components: {},
  data(){
    return {
      macClickNum: 0,
      lastLeft: 0,
      clickNum: 0,
      itemlist:[
        {id: 0, src: require("./assets/image1.jpg")},
        {id: 1, src: require("./assets/image2.jpg")},
        {id: 2, src: require("./assets/logo-Stanford_University-2020_08_11_14_58_29.png.jpg")},
        {id: 3, src: require("./assets/Ferrari F1-75.jpg")},
        {id: 4, src: require("./assets/image5.jpg")},
      ],
    };
  },
  methods: {
    leftScroll(){
      if(this.clickNum > 0){
        console.log(document.querySelectorAll(".adv-pad-item"));
        let width =
          document.querySelectorAll(".adv-pad-item")[this.clickNum - 1]
            .offsetWidth;
          // 公示：滚动距离（元素的magin-left值） = 它自己的长度 + 上一次滑动的距离
        console.log(document.getElementsByClassName("adv-pad"));
        document.getElementsByClassName("adv-pad")[0].style.marginLeft = `${
          this.lastLeft + width
        }px`;
        this.lastLeft = width + this.lastLeft;
          // 点击次数-3
        this.clickNum = this.clickNum - 1;
          // 如果点击次数小于最大点击次数，说明最后一个元素已经不在可是区域内了，显示右箭头
        if (this.clickNum < this.maxClickNum) {
          this.showRightIcon = true;
        }
      }
    },
    rightScroll(){
      if (this.clickNum < this.itemlist.length - 1) {
          // 获取当前元素宽度
        let width =
          document.querySelectorAll(".adv-pad-item")[this.clickNum].offsetWidth;
          // 获取最后一个元素距离左侧的距离
        let lastItemOffsetLeft =
          document.getElementsByClassName("adv-pad-item")[
          this.itemlist.length - 1
            ].offsetLeft;
          // 获取可视区域宽度
        const lookWidth = document.getElementById("adv-pad-scroll").clientWidth;
          // 如果最后一个元素距离左侧的距离大于可视区域的宽度，表示最后一个元素没有出现，执行滚动效果
        if (lastItemOffsetLeft > lookWidth) {
            // 公示：滚动距离（元素的magin-left值） = 负的它自己的长度 + 上一次滑动的距离
          document.getElementsByClassName("adv-pad")[0].style.marginLeft = `${
            -width + this.lastLeft
          }px`;
          this.lastLeft = -width + this.lastLeft;
            // 点击次数+3
          this.clickNum += 1;
            // 记录到最后一个元素出现在可视区域时，点击次数的最大值。用于后面点击左侧箭头时判断右侧箭头的显示
          this.maxClickNum = this.clickNum;
        }
        this.showRightIcon = lastItemOffsetLeft > lookWidth + width;
      }
    },  
  },
};
</script> -->




<template>
  <div>
    <h1 class="green">Welcome!</h1>
    <p class="green">This is a test web page</p>
    <div>
        <img id="pic" width="80%" :src="imgArr[0]">
    </div>

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
      var msg = document.getElementById("pic");
      if (i - 1 < 0)
        i = 4;
      msg.src = data.imgArr[i];
    },
    NextImage(){
      var msg = document.getElementById("pic");
      if (i + 1 > 4)
        i = 0;
      msg.src = data.imgArr[i];
    }
  },
}
</script>

<style scoped>
h1{
    font-weight: 500;
    font-size: 2.6rem;
    top: -10px;
}
p{
  font-weight: 500;
  font-size: 1.8rem;
  top: -10px;
}
</style>