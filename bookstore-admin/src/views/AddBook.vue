<template>
  <div class="add">
    <el-card class="add-container">
      <el-form :model="state.bookForm" :rules="state.rules" ref="bookRef" label-width="100px" class="bookForm">
        <el-form-item label="图书分类">
          <el-cascader :placeholder="state.defaultCate" style="width: 300px" :props="state.category"
                       @change="handleChangeCate"></el-cascader>
        </el-form-item>
        <el-form-item label="图书名称" prop="booksName">
          <el-input style="width: 300px" v-model="state.bookForm.booksName" placeholder="请输入图书名称"></el-input>
        </el-form-item>
        <el-form-item label="图书作者" prop="booksAuthor">
          <el-input style="width: 300px" v-model="state.bookForm.booksAuthor" placeholder="请输入图书作者"></el-input>
        </el-form-item>
        <el-form-item label="图书作者国籍" prop="booksAuthorCountry">
          <el-input style="width: 300px" v-model="state.bookForm.booksAuthorCountry" placeholder="请输入图书作者国籍"></el-input>
        </el-form-item>
        <el-form-item label="出版社出版时间" prop="booksPublish">
          <el-input style="width: 300px" v-model="state.bookForm.booksPublish" placeholder="出版社出版时间"></el-input>
        </el-form-item>
        <el-form-item label="图书简介" prop="booksIntro">
          <el-input style="width: 500px" type="textarea" v-model="state.bookForm.booksIntro"
                    placeholder="请输入图书简介(1000字)"></el-input>
        </el-form-item>
        <el-form-item label="图书价格" prop="originalPrice">
          <el-input type="number" min="0" style="width: 300px" v-model.number="state.bookForm.originalPrice"
                    placeholder="请输入图书价格"></el-input>
        </el-form-item>
        <el-form-item label="图书售卖价" prop="sellingPrice">
          <el-input type="number" min="0" style="width: 300px" v-model.number="state.bookForm.sellingPrice"
                    placeholder="请输入图书售价"></el-input>
        </el-form-item>
        <el-form-item label="图书库存" prop="stockNum">
          <el-input type="number" min="0" style="width: 300px" v-model.number="state.bookForm.stockNum"
                    placeholder="请输入图书库存"></el-input>
        </el-form-item>
        <el-form-item label="图书标签" prop="tag">
          <el-input style="width: 300px" v-model="state.bookForm.tag" placeholder="请输入图书小标签"></el-input>
        </el-form-item>
        <el-form-item label="上架状态" prop="booksSellStatus">
          <el-radio-group v-model.number="state.bookForm.booksSellStatus">
            <el-radio :label="0">上架</el-radio>
            <el-radio :label="1">下架</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item required label="图书主图" prop="booksCoverImg">
          <el-upload
              class="avatar-uploader"
              :action="state.uploadImgServer"
              accept="jpg,jpeg,png"
              :headers="{
              token: state.token
            }"
              :show-file-list="false"
              :before-upload="handleBeforeUpload"
              :on-success="handleUrlSuccess"
          >
            <img style="width: 100px; height: 100px; border: 1px solid #e9e9e9;" v-if="state.bookForm.booksCoverImg"
                 :src="state.bookForm.booksCoverImg" class="avatar">
            <el-icon v-else class="avatar-uploader-icon">
              <Plus/>
            </el-icon>
          </el-upload>
        </el-form-item>
        <el-form-item label="详情内容">
          <div ref='editor'></div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitAdd()">{{ state.id ? '立即修改' : '立即创建' }}</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import {reactive, ref, onMounted, onBeforeUnmount, getCurrentInstance} from 'vue'
import WangEditor from 'wangeditor'
import axios from '@/utils/axios'
import {ElMessage} from 'element-plus'
import {useRoute, useRouter} from 'vue-router'
import {localGet, uploadImgServer, uploadImgsServer} from '@/utils'

const {proxy} = getCurrentInstance()
const editor = ref(null)
const bookRef = ref(null)
const route = useRoute()
const router = useRouter()
const {id} = route.query
const state = reactive({
  uploadImgServer,
  token: localGet('token') || '',
  id: id,
  defaultCate: '',
  bookForm: {
    booksName: '',
    booksAuthor: '',
    booksAuthorCountry: '',
    booksPublish: '',
    booksIntro: '',
    originalPrice: 0,
    sellingPrice: 0,
    stockNum: 0,
    booksSellStatus: 0,
    booksCoverImg: '',
    tag: ''
  },
  rules: {
    booksName: [
      {required: 'true', message: '请填写图书名称', trigger: ['change']}
    ],
    originalPrice: [
      {required: 'true', message: '请填写图书价格', trigger: ['change']}
    ],
    sellingPrice: [
      {required: 'true', message: '请填写图书售价', trigger: ['change']}
    ],
    stockNum: [
      {required: 'true', message: '请填写图书库存', trigger: ['change']}
    ],
  },
  categoryId: '',
  category: {
    lazy: true,
    lazyLoad(node, resolve) {
      const {level = 0, value} = node
      axios.get('/categories', {
        params: {
          pageNumber: 1,
          pageSize: 1000,
          categoryLevel: level + 1,
          parentId: value || 0
        }
      }).then(res => {
        const list = res.list
        const nodes = list.map(item => ({
          value: item.categoryId,
          label: item.categoryName,
          leaf: level > 1
        }))
        resolve(nodes)
      })
    }
  }
})
let instance
onMounted(() => {
  instance = new WangEditor(editor.value)
  instance.config.showLinkImg = false
  instance.config.showLinkImgAlt = false
  instance.config.showLinkImgHref = false
  instance.config.uploadImgMaxSize = 2 * 1024 * 1024 // 2M
  instance.config.uploadFileName = 'file'
  instance.config.uploadImgHeaders = {
    token: state.token
  }
  // 图片返回格式不同，需要自定义返回格式
  instance.config.uploadImgHooks = {
    // 图片上传并返回了结果，想要自己把图片插入到编辑器中
    // 例如服务器端返回的不是 { errno: 0, data: [...] } 这种格式，可使用 customInsert
    customInsert: function (insertImgFn, result) {
      console.log('result', result)
      // result 即服务端返回的接口
      // insertImgFn 可把图片插入到编辑器，传入图片 src ，执行函数即可
      if (result.data && result.data.length) {
        result.data.forEach(item => insertImgFn(item))
      }
    }
  }
  instance.config.uploadImgServer = uploadImgsServer
  Object.assign(instance.config, {
    onchange() {
      console.log('change')
    },
  })
  instance.create()
  if (id) {
    axios.get(`/books/${id}`).then(res => {
      const {books, firstCategory, secondCategory, thirdCategory} = res
      state.bookForm = {
        booksName: books.booksName,
        booksAuthor: books.booksAuthor,
        booksAuthorCountry: books.booksAuthorCountry,
        booksPublish: books.booksPublish,
        booksIntro: books.booksIntro,
        originalPrice: books.originalPrice,
        sellingPrice: books.sellingPrice,
        stockNum: books.stockNum,
        booksSellStatus: books.booksSellStatus,
        booksCoverImg: proxy.$filters.prefix(books.booksCoverImg),
        tag: books.tag
      }
      state.categoryId = books.booksCategoryId
      state.defaultCate = `${firstCategory.categoryName}/${secondCategory.categoryName}/${thirdCategory.categoryName}`
      if (instance) {
        // 初始化图书详情 html
        instance.txt.html(books.booksDetailContent)
      }
    })
  }
})
onBeforeUnmount(() => {
  instance.destroy()
  instance = null
})
const submitAdd = () => {
  bookRef.value.validate((vaild) => {
    if (vaild) {
      // 默认新增用 post 方法
      let httpOption = axios.post
      let params = {
        booksCategoryId: state.categoryId,
        booksCoverImg: state.bookForm.booksCoverImg,
        booksDetailContent: instance.txt.html(),
        booksIntro: state.bookForm.booksIntro,
        booksName: state.bookForm.booksName,
        booksAuthor: state.bookForm.booksAuthor,
        booksAuthorCountry: state.bookForm.booksAuthorCountry,
        booksPublish: state.bookForm.booksPublish,
        booksSellStatus: state.bookForm.booksSellStatus,
        originalPrice: state.bookForm.originalPrice,
        sellingPrice: state.bookForm.sellingPrice,
        stockNum: state.bookForm.stockNum,
        tag: state.bookForm.tag
      }
      console.log('params', params)
      if (id) {
        params.booksId = id
        // 修改图书使用 put 方法
        httpOption = axios.put
      }
      httpOption('/books', params).then(() => {
        ElMessage.success(id ? '修改成功' : '添加成功')
        router.push({path: '/book'})
      })
    }
  })
}
const handleBeforeUpload = (file) => {
  const sufix = file.name.split('.')[1] || ''
  if (!['jpg', 'jpeg', 'png'].includes(sufix)) {
    ElMessage.error('请上传 jpg、jpeg、png 格式的图片')
    return false
  }
}
const handleUrlSuccess = (val) => {
  state.bookForm.booksCoverImg = val.data || ''
}
const handleChangeCate = (val) => {
  state.categoryId = val[2] || 0
}
</script>

<style scoped>
.add {
  display: flex;
}

.add-container {
  flex: 1;
  height: 100%;
}

.avatar-uploader {
  width: 100px;
  height: 100px;
  color: #ddd;
  font-size: 30px;
}

.avatar-uploader-icon {
  display: block;
  width: 100%;
  height: 100%;
  border: 1px solid #e9e9e9;
  padding: 32px 17px;
}
</style>