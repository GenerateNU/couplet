/* eslint-disable */
import React, { useEffect, useState } from "react";
import { StyleSheet, View } from "react-native";
import DropDownPicker from "react-native-dropdown-picker";
import scaleStyleSheet from "../../scaleStyles";
import { pronouns } from "../../utils/pronouns";

interface DropDownPronounProps {
  onPronounChange: (local: string) => void;
  selectedPronoun: string;
}

function DropDownPronoun({ onPronounChange, selectedPronoun }: DropDownPronounProps) {
  const [open, setOpen] = useState(false);
  const [value, setValue] = useState(selectedPronoun);
  const items = pronouns.map((neighborhood) => ({
    label: neighborhood,
    value: neighborhood
  }));

  useEffect(() => {
    onPronounChange(value);
  }, [value]);

  return (
    <View style={scaledStyles.container}>
      <DropDownPicker
        open={open}
        value={value}
        items={items}
        setOpen={setOpen}
        setValue={setValue}
        placeholder="Select a neighborhood"
      />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1
  }
});

const scaledStyles = scaleStyleSheet(styles);

export default DropDownPronoun;
