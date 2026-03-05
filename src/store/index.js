import { createStore } from "vuex";

export default createStore({
  state: {
    hospitals: ["Hospital A", "Hospital B", "Hospital C"],
    users: [
      {
        userid: "admin",
        password: "password",
        name: "Administrator",
        role: "admin",
        active: true,
      },
      {
        userid: "biller",
        password: "password",
        name: "Biller One",
        role: "biller",
        active: true,
      },
      {
        userid: "manager",
        password: "password",
        name: "Manager One",
        role: "manager",
        active: true,
      },
    ],
    currentUser: null,
    logs: [],
    medicines: [], // { name: 'Dolo', batches: [{ batchId: 'B1', ...}] }
    savedBills: [],
    currentBillNumber: 1,
  },
  getters: {
    getHospitals: (state) => state.hospitals,
    currentUser: (state) => state.currentUser,
    isAuthenticated: (state) => !!state.currentUser,
    userRole: (state) => (state.currentUser ? state.currentUser.role : null),
    allUsers: (state) => state.users,
    allLogs: (state) => state.logs,
    allMedicines: (state) => state.medicines,
    allBills: (state) => state.savedBills,
    nextBillNumber: (state) =>
      `BILL-${String(state.currentBillNumber).padStart(3, "0")}`,
    getBillById: (state) => (id) =>
      state.savedBills.find((b) => b.id == id || b.billNumber == id),
  },
  mutations: {
    SAVE_BILL(state, billData) {
      // Auto-increment bill number for next time
      state.currentBillNumber++;

      // Add metadata
      const newBill = {
        ...billData,
        id: Date.now(), // distinct from billNumber, used for routing
        timestamp: new Date().toLocaleString(),
        createdBy: state.currentUser ? state.currentUser.userid : "unknown",
        modifiedBy: null,
        isModified: false,
      };
      state.savedBills.unshift(newBill);
    },
    UPDATE_BILL(state, { billId, updatedData }) {
      const index = state.savedBills.findIndex(
        (b) => b.id == billId || b.billNumber == billId
      );
      if (index !== -1) {
        const original = state.savedBills[index];
        state.savedBills.splice(index, 1, {
          ...original,
          ...updatedData,
          modifiedBy: state.currentUser ? state.currentUser.userid : "unknown",
          isModified: true,
          lastModifiedAt: new Date().toLocaleString(),
        });
      }
    },
    ADD_HOSPITAL(state, hospital) {
      if (!state.hospitals.includes(hospital)) {
        state.hospitals.push(hospital);
      }
    },
    SET_USER(state, user) {
      state.currentUser = user;
    },
    ADD_LOG(state, log) {
      state.logs.unshift({
        ...log,
        id: Date.now(),
        timestamp: new Date().toLocaleString(),
      });
    },
    ADD_USER(state, user) {
      state.users.push({ ...user, active: true });
    },
    TOGGLE_USER_STATUS(state, userid) {
      const user = state.users.find((u) => u.userid === userid);
      if (user) {
        user.active = !user.active;
      }
    },
    ADD_MEDICINE_STOCK(state, { name, batchId, expiryDate, mrp, gst, qty }) {
      const existingMedicine = state.medicines.find(
        (m) => m.name.toLowerCase() === name.toLowerCase()
      );

      const batchData = {
        batchId,
        expiryDate,
        mrp: parseFloat(mrp),
        gst: parseFloat(gst),
        qty: parseInt(qty),
      };

      if (existingMedicine) {
        // check if batch exists
        const existingBatchIndex = existingMedicine.batches.findIndex(
          (b) => b.batchId === batchId
        );
        if (existingBatchIndex !== -1) {
          // Update existing batch (add qty)
          existingMedicine.batches[existingBatchIndex].qty += parseInt(qty);
          // Update other details in case they changed (last in wins)
          existingMedicine.batches[existingBatchIndex].expiryDate = expiryDate;
          existingMedicine.batches[existingBatchIndex].mrp = parseFloat(mrp);
          existingMedicine.batches[existingBatchIndex].gst = parseFloat(gst);
        } else {
          existingMedicine.batches.push(batchData);
        }
      } else {
        state.medicines.push({
          name: name,
          batches: [batchData],
        });
      }
    },
    DEDUCT_STOCK(state, { items }) {
      items.forEach((item) => {
        const medicine = state.medicines.find(
          (m) => m.name.toLowerCase() === item.medicineName.toLowerCase()
        );
        if (medicine) {
          const batch = medicine.batches.find(
            (b) => b.batchId === item.batchId
          );
          if (batch) {
            batch.qty -= item.quantity;
            if (batch.qty < 0) batch.qty = 0; // Prevent negative
          }
        }
      });
    },
    REMOVE_MEDICINE_BATCH(state, { name, batchId }) {
      const medicine = state.medicines.find((m) => m.name === name);
      if (medicine) {
        medicine.batches = medicine.batches.filter(
          (b) => b.batchId !== batchId
        );
        // Optional: Remove medicine if no batches left
        if (medicine.batches.length === 0) {
          state.medicines = state.medicines.filter((m) => m.name !== name);
        }
      }
    },
  },
  actions: {
    addHospital({ commit }, hospital) {
      commit("ADD_HOSPITAL", hospital);
    },
    login({ state, commit }, { userid, password }) {
      return new Promise((resolve, reject) => {
        // Simulate API delay
        setTimeout(() => {
          const user = state.users.find(
            (u) => u.userid === userid && u.password === password
          );
          if (user) {
            if (!user.active) {
              reject(new Error("Account is deactivated"));
              return;
            }
            commit("SET_USER", user);
            commit("ADD_LOG", {
              userid: user.userid,
              name: user.name,
              event: "Login",
            });
            resolve(user);
          } else {
            console.warn("Login failed for:", userid);
            reject(new Error("Invalid credentials"));
          }
        }, 500);
      });
    },
    logout({ state, commit }) {
      if (state.currentUser) {
        commit("ADD_LOG", {
          userid: state.currentUser.userid,
          name: state.currentUser.name,
          event: "Logout",
        });
      }
      commit("SET_USER", null);
    },
    createUser({ commit }, userData) {
      return new Promise((resolve) => {
        setTimeout(() => {
          commit("ADD_USER", userData);
          resolve(userData);
        }, 500);
      });
    },
    toggleUserStatus({ commit }, userid) {
      return new Promise((resolve) => {
        setTimeout(() => {
          commit("TOGGLE_USER_STATUS", userid);
          resolve();
        }, 300);
      });
    },
    addMedicineStock({ commit }, stockData) {
      return new Promise((resolve) => {
        setTimeout(() => {
          commit("ADD_MEDICINE_STOCK", stockData);
          resolve();
        }, 300);
      });
    },
    saveNewBill({ commit }, billData) {
      return new Promise((resolve) => {
        setTimeout(() => {
          commit("SAVE_BILL", billData);
          resolve();
        }, 300);
      });
    },
    updateBill({ commit }, payload) {
      return new Promise((resolve) => {
        setTimeout(() => {
          commit("UPDATE_BILL", payload);
          resolve();
        }, 300);
      });
    },
    deductStock({ commit }, items) {
      return new Promise((resolve) => {
        commit("DEDUCT_STOCK", { items });
        resolve();
      });
    },
    removeMedicineBatch({ commit }, payload) {
      return new Promise((resolve) => {
        setTimeout(() => {
          commit("REMOVE_MEDICINE_BATCH", payload);
          resolve();
        }, 300);
      });
    },
  },
  modules: {},
});
