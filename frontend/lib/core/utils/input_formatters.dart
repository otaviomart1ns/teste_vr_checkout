import 'package:flutter_multi_formatter/flutter_multi_formatter.dart';

class InputFormatters {
  static CurrencyInputFormatter currencyFormatter() {
    return CurrencyInputFormatter(
      leadingSymbol: '',
      useSymbolPadding: true,
      thousandSeparator: ThousandSeparator.Comma,
      mantissaLength: 2, //
    );
  }
}
