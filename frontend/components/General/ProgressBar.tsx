import React from 'react';
import { View, StyleSheet, Dimensions } from 'react-native';
import { COLORS } from './colors';

interface ProgressBarProps {
  totalCount: number;
  selectedCount: number;
}

const OnboardingBar: React.FC<ProgressBarProps> = ({ totalCount, selectedCount }) => {
  const totalWidth = totalCount * 62.8 + (totalCount - 1) * 8;
  const screenWidth = Dimensions.get('window').width;

  return (
    <View style={[styles.container, {justifyContent: 'center'}]}>
      <View style={{ flexDirection: 'row', width: totalWidth }}>
        {Array.from({ length: totalCount }).map((_, index) => (
          <View
            key={index}
            style={[
              styles.segment,
              index === selectedCount - 1 ? styles.segmentSelected : styles.segmentUnselected,
              { marginRight: index < totalCount - 1 ? 8 : 0 }, // Apply marginRight to all but the last segment
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
    width: 62.8,
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

export default OnboardingBar;
