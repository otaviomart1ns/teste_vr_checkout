import 'package:flutter_test/flutter_test.dart';
import 'package:frontend/modules/transaction/shared/models/local_transaction.dart';

void main() {
  test('copyWith deve manter e substituir campos corretamente', () {
    final original = LocalTransaction(
      id: 'a664d78d-cce6-4770-b287-b176a9e6e62a',
      description: 'Moto',
      date: DateTime(2020, 1, 2),
      amountUsd: 542.96,
    );

    final modificado = original.copyWith(
      description: 'Notebook',
      amountUsd: 799.99,
    );

    expect(modificado.id, equals('a664d78d-cce6-4770-b287-b176a9e6e62a'));
    expect(modificado.description, equals('Notebook'));
    expect(modificado.amountUsd, equals(799.99));
    expect(modificado.date, equals(DateTime(2020, 1, 2)));
  });
}
