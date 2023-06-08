import axios from '../utils/axios'

export function getDetail(id) {
  return axios.get(`/books/detail/${id}`);
}

export function getCategory() {
  return axios.get('/categories');
}

export function search(params) {
  return axios.get('/search', { params });
}

export function addComment(params) {
  return axios.post('/books/comment', params);
}

export function addCommentReply(params) {
  return axios.post('/books/commentReply', params);
}