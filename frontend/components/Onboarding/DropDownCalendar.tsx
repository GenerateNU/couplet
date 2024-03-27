import React, { useEffect, useState } from "react";
import { StyleSheet, View } from "react-native";
import DropDownPicker from "react-native-dropdown-picker";
import scaleStyleSheet from "../../scaleStyles";

function DropDownCalendar() {
  const [openDay, setOpenDay] = useState(false);
  const [openMonth, setOpenMonth] = useState(false);
  const [openYear, setOpenYear] = useState(false);

  const [day, setDay] = useState(1);
  const [month, setMonth] = useState(1);
  const [year, setYear] = useState(new Date().getFullYear());

  const days = Array.from({ length: 31 }, (_, i) => ({ label: `${i + 1}`, value: i + 1 }));
  const months = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December"
  ].map((month, index) => ({ label: month, value: index + 1 }));
  const currentYear = new Date().getFullYear();
  const years = Array.from({ length: 50 }, (_, i) => ({
    label: `${currentYear - i}`,
    value: currentYear - i
  }));

  return (
    <View style={scaledStyles.container}>
      <DropDownPicker
        open={openDay}
        value={day}
        items={days}
        setOpen={setOpenDay}
        setValue={setDay}
        containerStyle={scaledStyles.dropdown}
      />
      <DropDownPicker
        open={openMonth}
        value={month}
        items={months}
        setOpen={setOpenMonth}
        setValue={setMonth}
        containerStyle={scaledStyles.dropdown}
      />
      <DropDownPicker
        open={openYear}
        value={year}
        items={years}
        setOpen={setOpenYear}
        setValue={setYear}
        containerStyle={scaledStyles.dropdown}
      />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flexDirection: "row",
  },
  dropdown: {
    flex: 1,
    marginRight : 5,
  }
});

const scaledStyles = scaleStyleSheet(styles);

export default DropDownCalendar;
