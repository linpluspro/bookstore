<template>
  <div class="product-comment">
    <s-header :name="'图书评论'"></s-header>
    <div class="detail-content">
      <div class="detail-swipe-wrap">
        <van-swipe class="my-swipe" indicator-color="#1baeae">
          <van-swipe-item v-for="(item, index) in state.detail.booksCarouselList" :key="index">
            <img :src="item" alt="">
          </van-swipe-item>
        </van-swipe>
      </div>
      <div class="product-info">
        <div class="product-title">
          {{ state.detail.booksName || '' }}
        </div>
        <div class="product-desc">免邮费 顺丰快递</div>
        <div class="product-price">
          <span>¥{{ state.detail.sellingPrice || '' }}</span>
        </div>
      </div>

      <div v-clickoutside="hideReplyBtn" @click="inputFocus" class="my-reply">
        <div class="reply-info">
          <div
              tabindex="0"
              contenteditable="true"
              id="replyInput"
              spellcheck="false"
              placeholder="输入评论..."
              class="reply-input"
              v-click-outside="onClickOutside"
              @focus="showReplyBtn"
              @input="onDivInput($event)"
          >
          </div>
        </div>
        <div class="reply-btn-box" v-show="state.isShowCommentBtn">
          <el-button class="reply-btn" size="default" @click="sendComment" type="primary">发表评论</el-button>
        </div>
      </div>
    </div>

    <div class="comment-content">
      <div v-for="(item,i) in state.detail.comments" :key="i" class="author-title reply-father">
        <div class="author-info">
          <span class="author-name">{{ item.name }}</span>
          <span class="author-time">{{ item.time }}</span>
        </div>
        <div class="talk-box">
          <p>
            <span class="reply">{{ item.comment }}</span>
          </p>
        </div>
        <div class="reply-box">
          <div v-for="(reply,j) in item.reply" :key="j" class="author-title">
            <div class="author-info">
              <span class="author-name">{{ reply.name }}</span>
              <span class="author-time">{{ reply.time }}</span>
            </div>
            <div class="talk-box">
              <p>
                <span>回复 {{ reply.to }}:</span>
                <span class="reply">{{ reply.comment }}</span>
              </p>
            </div>

          </div>
        </div>
        <div v-show="_inputShow(i)" class="my-reply my-comment-reply">
          <div class="reply-info">
            <div tabindex="0" contenteditable="true" spellcheck="false" placeholder="输入评论..."
                 @input="onDivInput($event)"
                 class="reply-input reply-comment-input"></div>
          </div>
          <div class=" reply-btn-box">
            <el-button class="reply-btn" size="default" @click="sendCommentReply(i,j)" type="primary">发表评论</el-button>
          </div>
        </div>
      </div>
    </div>

    <van-action-bar>
      <van-action-bar-icon icon="cart-o" :badge="!cart.count ? '' : cart.count" @click="goTo()" text="购物车"/>
      <van-action-bar-button type="warning" @click="handleAddCart" text="加入购物车"/>
      <van-action-bar-button type="danger" @click="goToCart" text="立即购买"/>
    </van-action-bar>
  </div>
</template>

<script setup>
import {reactive, onMounted, nextTick} from 'vue'
import {useRoute, useRouter} from 'vue-router'
import {useCartStore} from '@/stores/cart'
import {getDetail, addComment, addCommentReply} from '@/service/book'
import {addCart} from '@/service/cart'
import sHeader from '@/components/SimpleHeader.vue'
import {showSuccessToast} from 'vant'
import {prefix} from '@/common/js/utils'
import {showFailToast} from "vant/lib/toast/function-call";
import {ClickOutside as vClickOutside} from 'element-plus';

const route = useRoute()
const router = useRouter()
const cart = useCartStore()

const state = reactive({
  detail: {
    booksId: 0,
    booksName: '',
    booksCarouselList: [],
    comments: [],
    isShowCommentBtn: false,
    replyComment: '',
  },
})

onMounted(async () => {
  const {id} = route.params
  const {data} = await getDetail(id)

  data.booksCarouselList = data.booksCarouselList.map(i => prefix(i))
  state.detail = data
  cart.updateCart()
})

nextTick(() => {
  // 一些和DOM有关的东西
  const content = document.querySelector('.detail-content')
  content.scrollTop = 0
})

// 点击外部关闭按钮
const onClickOutside = () => {
  state.isShowCommentBtn = false

  const replyInput = document.getElementById('replyInput');
  replyInput.style.padding = "10px"
  replyInput.style.border = "none"
}

function inputFocus() {
  const replyInput = document.getElementById('replyInput');
  replyInput.style.padding = "8px 8px"
  replyInput.style.border = "2px solid blue"
  replyInput.focus()
}

function showReplyBtn() {
  state.isShowCommentBtn = true
}

function _inputShow(i) {
  return state.detail.comments[i].inputShow
}

function hideReplyBtn() {
  state.isShowCommentBtn = false
  replyInput.style.padding = "10px"
  replyInput.style.border = "none"
}

function showReplyInput(i, name, id) {
  state.detail.comments[this.index].inputShow = false
  state.index = i
  state.detail.comments[i].inputShow = true
  state.to = name
  state.toId = id
}

function getCurrentDateTime() {
  const now = new Date();
  const isoDate = now.toISOString();
  return isoDate.replace(/Z$/, "+00:00");
}

const handleComment = async (params) => {
  const {resultCode} = await addComment(params)
  if (resultCode === 200) {
    showSuccessToast('发表成功！')
  } else {
    showFailToast("未购买该书，无权发表评论！")
  }
}

const handleCommentReply = async (params) => {
  const {resultCode} = await addCommentReply(params)
  if (resultCode === 200) {
    showSuccessToast('发表成功！')
  } else {
    showFailToast("未购买该书，无权发表评论！")
  }
}

function sendComment() {
  if (!state.replyComment) {
    showFailToast('评论不能为空!')
  } else {
    console.log(getCurrentDateTime())
    let params = {
      booksId: state.detail.booksId,
      comment: state.replyComment,
      time: getCurrentDateTime(),
      toId: -1,
    }
    document.getElementById("replyInput").innerHTML = ""
    handleComment(params)
  }
}

function sendCommentReply(i, j) {
  if (!state.replyComment) {
    showFailToast('评论不能为空!')
  } else {
    let params = {
      booksId: state.detail.booksId,
      comment: state.replyComment,
      time: getCurrentDateTime(),
      toId: state.detail.comments[i].id,
      name: state.detail.comments[i].name,
    }
    document.getElementsByClassName("reply-comment-input")[i].innerHTML = ""
    handleCommentReply(params)
  }
}

function onDivInput(e) {
  state.replyComment = e.target.innerHTML;
}

const goBack = () => {
  router.go(-1)
}

const goTo = () => {
  router.push({path: '/cart'})
}

const handleAddCart = async () => {
  const {resultCode} = await addCart({booksCount: 1, booksId: state.detail.booksId})
  if (resultCode === 200) showSuccessToast('添加成功')
  cart.updateCart()
}

const goToCart = async () => {
  await addCart({booksCount: 1, booksId: state.detail.booksId})
  cart.updateCart()
  router.push({path: '/cart'})
}

</script>

<style lang="less">
@import '../common/style/mixin';

.product-comment {
  .detail-header {
    position: fixed;
    top: 0;
    left: 0;
    z-index: 10000;
    .fj();
    .wh(100%, 44px);
    line-height: 44px;
    padding: 0 10px;
    .boxSizing();
    color: #252525;
    border-bottom: 1px solid #dcdcdc;

    .product-name {
      font-size: 14px;
    }
  }

  .detail-content {
    overflow: hidden;
    overflow-y: auto;

    .detail-swipe-wrap {
      .my-swipe .van-swipe-item {
        img {
          width: 100%;
        }
      }
    }

    .product-info {
      padding: 0 10px;

      .product-title {
        font-size: 18px;
        text-align: left;
      }

      .product-desc {
        font-size: 14px;
        text-align: left;
        color: #999;
        padding: 5px 0;
      }

      .product-price {
        .fj();

        span:nth-child(1) {
          color: #F63515;
          font-size: 22px;
        }

        span:nth-child(2) {
          color: #999;
          font-size: 16px;
        }
      }
    }
  }
}

.my-reply {
  padding: 10px;
  background-color: #fafbfc;
}

.my-reply .header-img {
  display: inline-block;
  vertical-align: top;
}

.my-reply .reply-info {
  display: inline-block;
  margin-left: 5px;
  width: 90%;
}

@media screen and (max-width: 1200px) {
  .my-reply .reply-info {
    width: 80%;
  }
}

.my-reply .reply-info .reply-input {
  min-height: 20px;
  line-height: 22px;
  padding: 10px 10px;
  color: #ccc;
  background-color: #fff;
  border-radius: 5px;
}

.my-reply .reply-info .reply-input:empty:before {
  content: attr(placeholder);
}

.my-reply .reply-info .reply-input:focus:before {
  content: none;
}

.my-reply .reply-info .reply-input:focus {
  padding: 8px 8px;
  border: 2px solid #00f;
  box-shadow: none;
  outline: none;
}

.my-reply .reply-btn-box {
  height: 25px;
  margin: 10px 0;
}

.my-reply .reply-btn-box .reply-btn {
  position: relative;
  float: right;
  margin-right: 15px;
}

.my-comment-reply {
  margin-left: 50px;
}

.my-comment-reply .reply-input {
  width: flex;
}

.author-title:not(:last-child) {
  border-bottom: 1px solid rgba(178, 186, 194, 0.3);
}

.author-title {
  padding: 10px;
}

.author-title .header-img {
  display: inline-block;
  vertical-align: top;
}

.author-title .author-info {
  display: inline-block;
  margin-left: 5px;
  width: 60%;
  height: 40px;
  line-height: 20px;
}

.author-title .author-info > span {
  display: block;
  cursor: pointer;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.author-title .author-info .author-name {
  color: #000;
  font-size: 18px;
  font-weight: bold;
}

.author-title .author-info .author-time {
  font-size: 14px;
}

.author-title .icon-btn {
  width: 30%;
  padding: 0 !important;
  float: right;
}

@media screen and (max-width: 1200px) {
  .author-title .icon-btn {
    width: 20%;
    padding: 7px;
  }
}

.author-title .icon-btn > span {
  cursor: pointer;
}

.author-title .icon-btn .iconfont {
  margin: 0 5px;
}

.author-title .talk-box {
  margin: 0 50px;
}

.author-title .talk-box > p {
  margin: 0;
}

.author-title .talk-box .reply {
  font-size: 16px;
  color: #000;
}

.author-title .reply-box {
  margin: 10px 0 0 50px;
  background-color: #efefef;
}

.comment-title {
  font-size: 18px;
  text-align: left;
}

.comment-content {
  height: calc(100vh - 50px);
}
</style>
