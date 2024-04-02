import { createSlice } from "@reduxjs/toolkit";

const formSlice = createSlice({
  name: "form",
  initialState: {
    name: ""
  },
  reducers: {
    setName: (state, action) => {
      state.name = action.payload;
    }
  }
});

export const { setName } = formSlice.actions;

export default formSlice.reducer;
