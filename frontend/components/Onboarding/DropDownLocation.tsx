/* eslint-disable */
import React, { useEffect, useState } from "react";
import { StyleSheet, View } from "react-native";
import DropDownPicker from "react-native-dropdown-picker";
import scaleStyleSheet from "../../scaleStyles";
import bostonNeighborhoods from "../../utils/location";

interface DropDownLocationProps {
  onLocationChange: (local: string) => void;
  selectedLocation: string;
}

function DropDownLocation({ onLocationChange, selectedLocation }: DropDownLocationProps) {
  const [open, setOpen] = useState(false);
  const [value, setValue] = useState(selectedLocation);
  const items = bostonNeighborhoods.map((neighborhood) => ({
    label: neighborhood,
    value: neighborhood
  }));

  useEffect(() => {
    onLocationChange(value);
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

export default DropDownLocation;
