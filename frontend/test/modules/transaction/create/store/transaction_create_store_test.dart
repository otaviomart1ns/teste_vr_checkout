import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
import 'package:frontend/modules/transaction/create/store/transaction_create_store.dart';
import 'package:frontend/modules/transaction/shared/transaction_service.dart';
import 'package:frontend/modules/transaction/shared/local_transaction_service.dart';
import 'package:frontend/modules/transaction/shared/models/local_transaction.dart';

class MockTransactionService extends Mock implements TransactionService {}

class MockLocalTransactionService extends Mock
    implements LocalTransactionService {}

void main() {
  late TransactionCreateStore store;
  late MockTransactionService mockTransactionService;
  late MockLocalTransactionService mockLocalService;

  setUpAll(() {
    registerFallbackValue(LocalTransaction(
      id: 'a664d78d-cce6-4770-b287-b176a9e6e62a',
      description: 'Moto',
      date: DateTime.parse('2020-01-02'),
      amountUsd: 542.96,
    ));
  });

  setUp(() {
    mockTransactionService = MockTransactionService();
    mockLocalService = MockLocalTransactionService();
    store = TransactionCreateStore(mockTransactionService, mockLocalService);
  });

  test('createTransactionLocal salva corretamente a transação', () async {
    when(() => mockLocalService.saveLocalTransaction(any()))
        .thenAnswer((_) async {});

    await store.createTransactionLocal(
      description: 'Moto',
      date: DateTime.parse('2020-01-02'),
      amountUsd: 542.96,
    );

    verify(() => mockLocalService.saveLocalTransaction(any())).called(1);
  });

  test('Estado inicial de isLoading é falso', () {
    expect(store.isLoading, isFalse);
  });

  test('Erro começa como null e pode ser setado manualmente', () {
    expect(store.errorMessage, isNull);
    store.errorMessage = 'Erro simulado';
    expect(store.errorMessage, 'Erro simulado');
  });
}
