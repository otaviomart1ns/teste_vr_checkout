import 'package:flutter_test/flutter_test.dart';
import 'package:frontend/core/utils/validators.dart';

void main() {
  group('validateDescription', () {
    test('retorna erro se vazio', () {
      expect(Validators.validateDescription(''), 'Descrição obrigatória');
    });

    test('retorna erro se tiver caracteres inválidos', () {
      expect(Validators.validateDescription('Moto@123'),
          'Somente letras e números');
    });

    test('retorna erro se for muito longa', () {
      expect(
          Validators.validateDescription('a' * 51), 'Máximo de 50 caracteres');
    });

    test('retorna null se válida', () {
      expect(Validators.validateDescription('Moto123'), isNull);
    });
  });

  group('validateValue', () {
    test('retorna erro se vazio', () {
      expect(Validators.validateValue(''), 'Informe um valor');
    });

    test('retorna erro se for zero ou negativo', () {
      expect(
          Validators.validateValue('-10'), 'Valor inválido (máximo 99.999,99)');
    });

    test('retorna erro se for texto', () {
      expect(
          Validators.validateValue('abc'), 'Valor inválido (máximo 99.999,99)');
    });

    test('retorna null se válido', () {
      expect(Validators.validateValue('542.96'), isNull);
    });
  });

  group('validateTransactionId', () {
    test('retorna erro se vazio', () {
      expect(Validators.validateTransactionId(''), 'Informe o ID da transação');
    });

    test('retorna erro se inválido', () {
      expect(Validators.validateTransactionId('1234'),
          'ID da transação inválido (UUID esperado)');
    });

    test('retorna null se UUID válido', () {
      expect(
          Validators.validateTransactionId(
              'a664d78d-cce6-4770-b287-b176a9e6e62a'),
          isNull);
    });
  });

  group('validateTransactionDate', () {
    test('retorna erro se null', () {
      expect(Validators.validateTransactionDate(null),
          'Data da transação obrigatória');
    });

    test('retorna erro se futura', () {
      final future = DateTime.now().add(const Duration(days: 1));
      expect(Validators.validateTransactionDate(future),
          'Data futura não permitida');
    });

    test('retorna erro se antiga demais', () {
      final old = DateTime.now().subtract(const Duration(days: 365 * 6));
      expect(Validators.validateTransactionDate(old),
          'Data muito antiga (máx. 5 ano)');
    });

    test('retorna null se data válida', () {
      final date = DateTime.now().subtract(const Duration(days: 365 * 3));
      expect(Validators.validateTransactionDate(date), isNull);
    });
  });
}
