import React from "react";
import { StyleSheet, View } from "react-native";
import scaleStyleSheet from "../../scaleStyles";
import ReactionButton from "./ReactionButton";

export type ReactionProps = {
  handleReact: (like: boolean) => void;
};

function Reaction({ handleReact }: ReactionProps) {
  return (
    <View style={scaledStyles.container}>
      <ReactionButton like={false} icon="window-close" handleReact={handleReact} />
      <ReactionButton like icon="heart" handleReact={handleReact} />
    </View>
  );
}

export default Reaction;

const styles = StyleSheet.create({
  container: {
    flexDirection: "row",
    justifyContent: "space-between",
    paddingHorizontal: 20
  }
});
const scaledStyles = scaleStyleSheet(styles);
