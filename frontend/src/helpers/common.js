export default class CommonHelper {
  static convertYen(src) {
    if (!src)
      return '-';
    let result = '';
    if (src > 10000) {
      result = `${this.round(src / 10000, 1)}万円`;
    } else {
      result = `${src}円`;
    }
    return result;
  }

  static round(value, precision) {
    // eslint-disable-next-line no-restricted-properties
    const multiplier = Math.pow(10, precision || 0);
    return Math.round(value * multiplier) / multiplier;
  }
}
