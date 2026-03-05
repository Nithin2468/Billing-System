import axios from "axios";

const apiClient = axios.create({
  baseURL: "http://localhost:8090", // Replace with your actual backend URL
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
});

export default {
  saveBill(billData) {
    return apiClient.post("/billingrecords", billData);
  },
  getBankDetails(ifsc) {
    return apiClient.get(`/get-bank-details?ifsc=${ifsc}`);
  },
  downloadPdf(htmlContent) {
    return apiClient.post(
      "/downloadpdff",
      { htmlContent },
      {
        responseType: "blob", // Important for handling binary PDF data
      }
    );
  },
  // Medicine Stock API
  addMedicine(stockData) {
    return apiClient.post("/add-medicine", stockData);
  },
  getMedicines() {
    return apiClient.get("/get-medicines");
  },
  updateMedicine(stockData) {
    return apiClient.post("/update-medicine", stockData);
  },
  // Billing History API
  getBills() {
    return apiClient.get("/get-bills");
  },
  getBillDetails(billNumber) {
    return apiClient.get(`/get-bill-details?billNumber=${billNumber}`);
  },

  // User Management API
  createUser(userData) {
    return apiClient.post("/newuser", userData);
  },
  getUsers() {
    return apiClient.get("/get-users");
  },
  updateUserStatus(userId, status) {
    return apiClient.post("/update-user-status", {
      user_id: userId,
      status: status,
    });
  },
  // Stock Report API
  getStockReport(params) {
    return apiClient.get("/reports/stock-sales", { params });
  },

  // Auth API
  login(credentials) {
    return apiClient.post("/loginApi", credentials);
  },
};
