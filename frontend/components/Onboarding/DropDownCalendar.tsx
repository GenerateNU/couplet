/* eslint-disable */
import React, { useEffect, useState } from "react";
import { StyleSheet, View } from "react-native";
import DropDownPicker from "react-native-dropdown-picker";
import scaleStyleSheet from "../../scaleStyles";
import { screenHeight } from "../../utils/dimensions";

interface DropDownCalendarProps {
  onDateChange: (day: number, month: number, year: number) => void;
  onDropDownOpen: (openDay: boolean, openMonth: boolean, openYear: boolean) => void;
}

function DropDownCalendar({ onDateChange, onDropDownOpen }: DropDownCalendarProps) {
  const [openDay, setOpenDay] = useState(false);
  const [openMonth, setOpenMonth] = useState(false);
  const [openYear, setOpenYear] = useState(false);

  const [day, setDay] = useState(0);
  const [month, setMonth] = useState(0);
  const [year, setYear] = useState(0);

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
  ].map((monthParam, index) => ({ label: monthParam, value: index + 1 }));
  const currentYear = new Date().getFullYear();
  const years = Array.from({ length: 50 }, (_, i) => ({
    label: `${currentYear - i}`,
    value: currentYear - i
  }));

  useEffect(() => {
    onDateChange(day, month, year);
  }, [day, month, year]);

  return (
    <View style={scaledStyles.container}>
      <DropDownPicker
        open={openDay}
        value={day}
        items={days}
        setOpen={setOpenDay}
        setValue={setDay}
        containerStyle={scaledStyles.dropdown}
        onOpen={() => onDropDownOpen(true, openMonth, openYear)}
        onClose={() => onDropDownOpen(false, openMonth, openYear)}
        dropDownContainerStyle={{ height: screenHeight * 0.15 }}
        placeholder="Day"
      />
      <DropDownPicker
        open={openMonth}
        value={month}
        items={months}
        setOpen={setOpenMonth}
        setValue={setMonth}
        containerStyle={scaledStyles.monthDropDown}
        onOpen={() => onDropDownOpen(openDay, true, openYear)}
        onClose={() => onDropDownOpen(openDay, false, openYear)}
        dropDownContainerStyle={{ height: screenHeight * 0.15 }}
        placeholder="Month"
      />
      <DropDownPicker
        open={openYear}
        value={year}
        items={years}
        setOpen={setOpenYear}
        setValue={setYear}
        containerStyle={scaledStyles.dropdown}
        // Can remove this
        onOpen={() => onDropDownOpen(openDay, openMonth, true)}
        onClose={() => onDropDownOpen(openDay, openMonth, false)}
        dropDownContainerStyle={{ height: screenHeight * 0.15 }}
        placeholder="Year"
      />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flexDirection: "row"
  },
  dropdown: {
    flex: 1,
    marginRight: 5
  },
  monthDropDown: {
    flex: 1.5,
    marginRight: 5
  }
});

const scaledStyles = scaleStyleSheet(styles);

export default DropDownCalendar;
