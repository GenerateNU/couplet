import React from 'react';
import { View, StyleSheet } from 'react-native';
import COLORS from '../../colors';
import scaleStyleSheet from '../../scaleStyles';
import { scaleWidth } from '../../ratio';

interface OnboardingBarProps {
  selectedCount: number;
}

function OnboardingBar({selectedCount}: OnboardingBarProps) {

  return (
    <View style={[scaledStyles.container, {justifyContent: 'center'}]}>
      <View style={{ flexDirection: 'row', width: scaleWidth(346) }}>
        {Array.from({ length: 4 }).map((_, index) => (
          <View
            style={[
              scaledStyles.segment,
              index === selectedCount - 1 ? scaledStyles.segmentSelected : scaledStyles.segmentUnselected,
              { marginRight: index < 3 ? 8 : 0 },
            ]}
          />
        ))}
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    width: '100%',
    height: 21,
    alignItems: 'center',
  },
  segment: {
    width: 80.5,
    height: 3,
    borderRadius: 100,
  },
  segmentSelected: {
    backgroundColor: COLORS.primary,
  },
  segmentUnselected: {
    backgroundColor: COLORS.disabled,
  },
});

const scaledStyles = scaleStyleSheet(styles);

export default OnboardingBar;
