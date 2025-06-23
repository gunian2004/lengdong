import { createRouter, createWebHistory } from "vue-router";
import AuthCard from "../components/AuthCard.vue";
import Factory from "../components/Factory.vue";
import Distributor from "../components/Distributor.vue";
import Retailer from "../components/Retailer.vue";
import Regulator from "../components/Regulator.vue";
import Consumer from "../components/Consumer.vue";
import DashBoard from "@/components/DashBoard.vue";
import BlockchainTest from "../components/BlockchainTest.vue";

const routes = [
  {
    path: "/dashboard",
    name: "DashBoard",
    component: DashBoard,
  },
  {
    path: "/",
    name: "AuthCard",
    component: AuthCard,
  },
  {
    path: "/factory",
    name: "Factory",
    component: Factory,
  },
  {
    path: "/distributor",
    name: "Distributor",
    component: Distributor,
  },
  {
    path: "/retailer",
    name: "retailer",
    component: Retailer,
  },
  {
    path: "/regulator",
    name: "Regulator",
    component: Regulator,
  },
  {
    path: "/consumer",
    name: "Consumer",
    component: Consumer,
  },
  {
    path: "/blockchain-test",
    name: "BlockchainTest",
    component: BlockchainTest,
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
