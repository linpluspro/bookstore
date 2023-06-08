<template>
  <div class="cart-box">
    <s-header :name="'购物车'" :noback="true"></s-header>
    <div class="cart-body">
      <van-checkbox-group @change="groupChange" v-model="state.result" ref="checkboxGroup">
        <van-swipe-cell :right-width="50" v-for="(item, index) in state.list" :key="index">
          <div class="book-item">
            <van-checkbox :name="item.cartItemId"/>
            <div class="book-img"><img :src="$filters.prefix(item.booksCoverImg)" alt=""></div>
            <div class="book-desc">
              <div class="book-title">
                <span>{{ item.booksName }}</span>
                <span>x{{ item.booksCount }}</span>
              </div>
              <div class="book-btn">
                <div class="price">¥{{ item.sellingPrice }}</div>
                <van-stepper
                    integer
                    :min="1"
                    :max="5"
                    :model-value="item.booksCount"
                    :name="item.cartItemId"
                    async-change
                    @change="onChange"
                />
              </div>
            </div>
          </div>
          <template #right>
            <van-button
                square
                icon="delete"
                type="danger"
                class="delete-button"
                @click="deleteBook(item.cartItemId)"
            />
          </template>
        </van-swipe-cell>
      </van-checkbox-group>
    </div>
    <van-submit-bar
        v-if="state.list.length > 0"
        class="submit-all van-hairline--top"
        :price="total * 100"
        button-text="结算"
        button-type="primary"
        @submit="onSubmit"
    >
      <van-checkbox @click="allCheck" v-model:checked="state.checkAll">全选</van-checkbox>
    </van-submit-bar>
    <div class="empty" v-if="!state.list.length">
      <img class="empty-cart" src="https://s.yezgea02.com/1604028375097/empty-car.png" alt="空购物车">
      <div class="title">购物车空空如也</div>
      <van-button round color="#1baeae" type="primary" @click="goTo" block>前往选购</van-button>
    </div>
    <nav-bar></nav-bar>
  </div>
</template>

<script setup>
import {reactive, onMounted, computed} from 'vue'
import {useRouter} from 'vue-router'
import {useCartStore} from '@/stores/cart'
import {showToast, showLoadingToast, closeToast, showFailToast} from 'vant'
import navBar from '@/components/NavBar.vue'
import sHeader from '@/components/SimpleHeader.vue'
import {getCart, deleteCartItem, modifyCart} from '@/service/cart'

const router = useRouter()
const cart = useCartStore()
const state = reactive({
  checked: false,
  list: [],
  all: false,
  result: [],
  checkAll: true
})

onMounted(() => {
  init()
})

const init = async () => {
  showLoadingToast({message: '加载中...', forbidClick: true});
  const {data} = await getCart({pageNumber: 1})
  state.list = data
  state.result = data.map(item => item.cartItemId)
  closeToast()
}

const total = computed(() => {
  let sum = 0
  let _list = state.list.filter(item => state.result.includes(item.cartItemId))
  _list.forEach(item => {
    sum += item.booksCount * item.sellingPrice
  })
  return sum
})

const goBack = () => {
  router.go(-1)
}

const goTo = () => {
  router.push({path: '/home'})
}

const onChange = async (value, detail) => {
  if (value > 5) {
    showFailToast('超出单个图书的最大购买数量')
    return
  }
  if (value < 1) {
    showFailToast('图书不得小于0')
    return
  }
  /**
   * 这里的操作是因为，后面修改购物车后，手动添加的计步器的数据，为了防止数据不对
   * 这边做一个拦截处理，如果点击的时候，购物车单项的 booksCount 等于点击的计步器数字，
   * 那么就不再进行修改操作
   */
  if (state.list.find(item => item.cartItemId === detail.name)?.booksCount === value) return
  showLoadingToast({message: '修改中...', forbidClick: true});
  const params = {
    cartItemId: detail.name,
    booksCount: value
  }
  await modifyCart(params)
  /**
   * 修改完成后，没有请求购物车列表，是因为闪烁的问题，
   * 这边手动给操作的购物车图书修改数据
   */
  state.list.forEach(item => {
    if (item.cartItemId === detail.name) {
      item.booksCount = value
    }
  })
  closeToast()
}

const onSubmit = async () => {
  if (state.result.length === 0) {
    showFailToast('请选择图书进行结算')
    return
  }
  const params = JSON.stringify(state.result)
  router.push({path: '/create-order', query: {cartItemIds: params}})
}

const deleteBook = async (id) => {
  await deleteCartItem(id)
  cart.updateCart()
  init()
}

const groupChange = (result) => {
  if (result.length === state.list.length) {
    state.checkAll = true
  } else {
    state.checkAll = false
  }
  state.result = result
}

const allCheck = () => {
  if (!state.checkAll) {
    state.result = state.list.map(item => item.cartItemId)
  } else {
    state.result = []
  }
}
</script>

<style lang="less">
@import '../common/style/mixin';

.cart-box {
  .cart-header {
    position: fixed;
    top: 0;
    left: 0;
    z-index: 10000;
    .fj();
    .wh(100%, 44px);
    line-height: 44px;
    padding: 0 10px;
    .boxSizing();
    border-bottom: 1px solid #dcdcdc;

    .cart-name {
      font-size: 14px;
    }
  }

  .cart-body {
    margin: 16px 0 100px 0;
    padding-left: 10px;

    .book-item {
      display: flex;

      .book-img {
        img {
          .wh(60px, 80px)
        }
      }

      .book-desc {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        flex: 1;
        padding: 20px;

        .book-title {
          font-size: 14px;
          display: flex;
          justify-content: space-between;
        }

        .book-btn {
          display: flex;
          justify-content: space-between;

          .price {
            font-size: 16px;
            color: red;
            line-height: 28px;
          }

          .van-icon-delete {
            font-size: 20px;
            margin-top: 4px;
          }
        }
      }
    }

    .delete-button {
      width: 50px;
      height: 100%;
    }
  }

  .empty {
    width: 50%;
    margin: 0 auto;
    text-align: center;
    margin-top: 200px;

    .empty-cart {
      width: 150px;
      margin-bottom: 20px;
    }

    .van-icon-smile-o {
      font-size: 50px;
    }

    .title {
      font-size: 16px;
      margin-bottom: 20px;
    }
  }

  .submit-all {
    margin-bottom: 64px;

    .van-checkbox {
      margin-left: 10px
    }

    .van-submit-bar__text {
      margin-right: 10px
    }

    .van-submit-bar__button {
      background: @primary;
    }
  }

  .van-checkbox__icon--checked .van-icon {
    background-color: @primary;
    border-color: @primary;
  }
}
</style>
