import React from "react";
import { StyleSheet, View } from "react-native";
import scaleStyleSheet from "../../scaleStyles";
import ReactionButton from "./ReactionButton";

export type ReactionProps = {
  handleReact: (like: boolean) => void;
};

function Reaction({ handleReact }: ReactionProps) {
  return (
<<<<<<< HEAD
    <View
      style={{
        flexDirection: "row",
        justifyContent: "space-between",
        paddingHorizontal: 20,
        paddingVertical: 10
      }}
    >
=======
    <View style={scaledStyles.container}>
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
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
