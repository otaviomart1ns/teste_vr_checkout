// test/core/utils/input_formatters_test.dart
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_multi_formatter/flutter_multi_formatter.dart';
import 'package:frontend/core/utils/input_formatters.dart';

void main() {
  group('InputFormatters.currencyFormatter', () {
    test('deve retornar um CurrencyInputFormatter configurado corretamente',
        () {
      final formatter = InputFormatters.currencyFormatter();
      expect(formatter, isA<CurrencyInputFormatter>());

      expect(formatter.thousandSeparator, ThousandSeparator.Comma);
      expect(formatter.mantissaLength, 2);
    });

    test('deve formatar corretamente um número simples', () {
      final formatter = InputFormatters.currencyFormatter();
      final result = formatter.formatEditUpdate(
        const TextEditingValue(text: '1234'),
        const TextEditingValue(text: '12345'),
      );

      expect(result.text, contains(','));
    });
  });
}
