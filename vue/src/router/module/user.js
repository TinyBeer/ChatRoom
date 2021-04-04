const userRoutes = [

  {
    path: '/register',
    name: 'register',
    // 惰性加载
    component: () => import('@/views/register/Register.vue'),
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/login/Login.vue'),
  },
];

export default userRoutes;
