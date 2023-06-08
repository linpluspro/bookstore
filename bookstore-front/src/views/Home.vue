<template>
  <div>
    <header class="home-header wrap" :class="{ 'active': state.headerScroll }">
      <router-link tag="i" to="./category"><i class="nbicon nbmenu2"></i></router-link>
      <div class="header-search">
        <span class="app-name">网上书店</span>
        <i class="iconfont icon-search"></i>
        <router-link tag="span" class="search-title" to="./product-list?from=home">书名、描述...</router-link>
      </div>
      <router-link class="login" tag="span" to="./login" v-if="!state.isLogin">登录</router-link>
      <router-link class="login" tag="span" to="./user" v-else>
        <van-icon name="manager-o" />
      </router-link>
    </header>
    <nav-bar />
    <swiper :list="state.swiperList"></swiper>
    <div class="category-list">
      <div v-for="item in state.categoryList" v-bind:key="item.categoryId" @click="tips">
        <img :src="item.imgUrl">
        <span>{{ item.name }}</span>
      </div>
    </div>
    <div class="book">
      <header class="book-header">新品上线</header>
      <van-skeleton title :row="3" :loading="state.loading">
        <div class="book-box">
          <div class="book-item" v-for="item in state.newBooks" :key="item.booksId" @click="goToDetail(item)">
            <img :src="$filters.prefix(item.booksCoverImg)" alt="" width="80" height="160">
            <div class="book-desc">
              <div class="title">{{ item.booksName }}</div>
              <div class="price">¥ {{ item.sellingPrice }}</div>
            </div>
          </div>
        </div>
      </van-skeleton>
    </div>
    <div class="book">
      <header class="book-header">热门图书</header>
      <van-skeleton title :row="3" :loading="state.loading">
        <div class="book-box">
          <div class="book-item" v-for="item in state.hots" :key="item.booksId" @click="goToDetail(item)">
            <img :src="$filters.prefix(item.booksCoverImg)" alt="" width="80" height="160">
            <div class="book-desc">
              <div class="title">{{ item.booksName }}</div>
              <div class="price">¥ {{ item.sellingPrice }}</div>
            </div>
          </div>
        </div>
      </van-skeleton>
    </div>
    <div class="book" :style="{ paddingBottom: '100px' }">
      <header class="book-header">最新推荐</header>
      <van-skeleton title :row="3" :loading="state.loading">
        <div class="book-box">
          <div class="book-item" v-for="item in state.recommends" :key="item.booksId" @click="goToDetail(item)">
            <img :src="$filters.prefix(item.booksCoverImg)" alt="" width="80" height="160">
            <div class="book-desc">
              <div class="title">{{ item.booksName }}</div>
              <div class="price">¥ {{ item.sellingPrice }}</div>
            </div>
          </div>
        </div>
      </van-skeleton>
    </div>
  </div>
</template>

<script setup>
import { reactive, onMounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import swiper from '@/components/Swiper.vue'
import navBar from '@/components/NavBar.vue'
import { getHome } from '@/service/home'
import { getLocal } from '@/common/js/utils'
import { showLoadingToast, closeToast, showToast } from 'vant'
import { useCartStore } from '@/stores/cart'

const cart = useCartStore()
const router = useRouter()
const state = reactive({
  swiperList: [], // 轮播图列表
  isLogin: false, // 是否已登录
  headerScroll: false, // 滚动透明判断
  hots: [],
  newBooks: [],
  recommends: [],
  categoryList: [
    {
      name: '全球购',
      imgUrl: 'https://s.yezgea02.com/1604041127880/%E5%85%A8%E7%90%83%E8%B4%AD%402x.png',
      categoryId: 100002
    }, {
      name: '充值缴费',
      imgUrl: 'https://s.yezgea02.com/1604041127880/%E5%85%85%E5%80%BC%402x.png',
      categoryId: 100006
    }, {
      name: '领劵',
      imgUrl: 'https://s.yezgea02.com/1604041127880/%E9%A2%86%E5%88%B8%402x.png',
      categoryId: 100008
    }, {
      name: '省钱',
      imgUrl: 'https://s.yezgea02.com/1604041127880/%E7%9C%81%E9%92%B1%402x.png',
      categoryId: 100009
    }, {
      name: '全部',
      imgUrl: 'https://s.yezgea02.com/1604041127880/%E5%85%A8%E9%83%A8%402x.png',
      categoryId: 100010
    }
  ],
  loading: true
})
onMounted(async () => {
  const token = getLocal('token')
  if (token) {
    state.isLogin = true
    // 获取购物车数据.
    cart.updateCart()
  }
  showLoadingToast({
    message: '加载中...',
    forbidClick: true
  });
  const { data } = await getHome()
  state.swiperList = data.carousels
  state.newBooks = data.newBooks
  state.hots = data.hotBooks
  state.recommends = data.recommendBooks
  state.loading = false
  closeToast()
})

nextTick(() => {
  document.body.addEventListener('scroll', () => {
    let scrollTop = window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop
    scrollTop > 100 ? state.headerScroll = true : state.headerScroll = false
  })
})

const goToDetail = (item) => {
  router.push({ path: `/product/${item.booksId}` })
}

const tips = () => {
  showToast('敬请期待');
}
</script>

<style lang="less" scoped>
@import '../common/style/mixin';

.home-header {
  position: fixed;
  left: 0;
  top: 0;
  .wh(100%, 50px);
  .fj();
  line-height: 50px;
  padding: 0 15px;
  .boxSizing();
  font-size: 15px;
  color: #fff;
  z-index: 10000;

  .nbmenu2 {
    color: @primary;
  }

  &.active {
    background: @primary;

    .nbmenu2 {
      color: #fff;
    }

    .login {
      color: #fff;
    }
  }

  .header-search {
    display: flex;
    width: 74%;
    line-height: 20px;
    margin: 10px 0;
    padding: 5px 0;
    color: #232326;
    background: rgba(255, 255, 255, .7);
    border-radius: 20px;

    .app-name {
      padding: 0 10px;
      color: @primary;
      font-size: 20px;
      font-weight: bold;
      border-right: 1px solid #666;
    }

    .icon-search {
      padding: 0 10px;
      font-size: 17px;
    }

    .search-title {
      font-size: 12px;
      color: #666;
      line-height: 21px;
    }
  }

  .icon-iconyonghu {
    color: #fff;
    font-size: 22px;
  }

  .login {
    color: @primary;
    line-height: 52px;

    .van-icon-manager-o {
      font-size: 20px;
      vertical-align: -3px;
    }
  }
}

.category-list {
  background: #fff;
  display: flex;
  flex-shrink: 0;
  flex-wrap: wrap;
  width: 100%;
  padding-bottom: 13px;

  div {
    display: flex;
    flex-direction: column;
    width: 20%;
    text-align: center;

    img {
      .wh(36px, 36px);
      margin: 13px auto 8px auto;
    }
  }
}

.book {
  .book-header {
    background: #f9f9f9;
    height: 50px;
    line-height: 50px;
    text-align: center;
    color: @primary;
    font-size: 16px;
    font-weight: 500;
  }

  .book-box {
    background: #fff;
    display: flex;
    justify-content: flex-start;
    flex-wrap: wrap;

    .book-item {
      box-sizing: border-box;
      width: 50%;
      border-bottom: 2PX solid #e9e9e9;
      padding: 10px 10px;

      img {
        display: block;
        width: 120px;
        margin: 0 auto;
      }

      .book-desc {
        text-align: center;
        font-size: 14px;
        padding: 10px 0;

        .title {
          color: @primary;
        }

        .price {
          color: @primary;
        }
      }

      &:nth-child(2n + 1) {
        border-right: 1PX solid #e9e9e9;
      }
    }
  }
}

.floor-list {
  width: 100%;
  padding-bottom: 50px;

  .floor-head {
    width: 100%;
    height: 40px;
    background: #F6F6F6;
  }

  .floor-content {
    display: flex;
    flex-shrink: 0;
    flex-wrap: wrap;
    width: 100%;
    .boxSizing();

    .floor-category {
      width: 50%;
      padding: 10px;
      border-right: 1px solid #dcdcdc;
      border-bottom: 1px solid #dcdcdc;
      .boxSizing();

      &:nth-child(2n) {
        border-right: none;
      }

      p {
        font-size: 17px;
        color: #333;

        &:nth-child(2) {
          padding: 5px 0;
          font-size: 13px;
          color: @primary;
        }
      }

      .floor-products {
        .fj();
        width: 100%;

        img {
          .wh(65px, 65px);
        }
      }
    }
  }
}
</style>