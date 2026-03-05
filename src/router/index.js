import { createRouter, createWebHashHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import BillingView from "@/views/BillingView.vue";
import LoginView from "../views/LoginView.vue";
import DashboardView from "../views/DashboardView.vue";
import LogView from "../views/LogView.vue";
import UsersView from "../views/UsersView.vue";

const routes = [
  {
    path: "/",
    name: "home",
    component: HomeView,
  },

  {
    path: "/login",
    name: "login",
    component: LoginView,
  },
  {
    path: "/dashboard",
    component: DashboardView,
    meta: { requiresAuth: true },
    children: [
      {
        path: "",
        name: "dashboard-home",
        component: () => import("../components/DashboardPage.vue"),
      },
      {
        path: "billing",
        name: "billing",
        component: BillingView,
      },
      {
        path: "billing/:id",
        name: "billing-edit",
        component: BillingView,
        meta: { roles: ["admin", "manager"] }, // Only managers can edit?
      },
      {
        path: "history",
        name: "history",
        component: () => import("../components/BillingHistory.vue"),
        meta: { roles: ["admin", "manager"] },
      },
      {
        path: "stocks",
        name: "stocks",
        component: () => import("../views/MedicineStockView.vue"),
        meta: { roles: ["admin", "manager"] },
      },
      {
        path: "available-medicines",
        name: "available-medicines",
        component: () => import("../views/AvailableMedicineView.vue"),
        // No meta role restriction -> accessible to all Authenticated
      },
      {
        path: "reports/stock",
        name: "stock-report",
        component: () => import("../views/ManagerStockReport.vue"),
        meta: { roles: ["admin", "manager"] },
      },
      {
        path: "logs",
        name: "logs",
        component: LogView,
        meta: { roles: ["admin", "manager"] },
      },
      {
        path: "users",
        name: "users",
        component: UsersView,
        meta: { roles: ["admin", "manager"] },
      },
    ],
  },
];

import store from "@/store";

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const isAuthenticated = store.getters.isAuthenticated;
  const userRole = store.getters.userRole;

  if (to.meta.requiresAuth && !isAuthenticated) {
    next({ name: "login" });
  } else if (
    to.meta.roles &&
    (!isAuthenticated || !to.meta.roles.includes(userRole))
  ) {
    next({ name: "dashboard-home" }); // Redirect to dashboard if role not allowed
  } else {
    next();
  }
});

export default router;
