import request from '@/utils/request';

// 用户注册
const register = ({ name, id, password }) => {
  return request.post('auth/register', { name, id, password });
};

// 获取信息
const info = () => {
  return request.get('auth/info');
};

// 用户登录
const login = ({ id, password }) => {
  return request.post('auth/login', { id, password });
};

export default {
  register,
  info,
  login,
};
