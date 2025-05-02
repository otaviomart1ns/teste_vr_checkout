import 'package:flutter_test/flutter_test.dart';
import 'package:frontend/modules/transaction/shared/local_transaction_service.dart';
import 'package:frontend/modules/transaction/shared/models/local_transaction.dart';
import 'package:hive/hive.dart';
import 'package:hive_test/hive_test.dart';

void main() {
  late LocalTransactionService service;

  final fakeTransaction = LocalTransaction(
    id: 'a664d78d-cce6-4770-b287-b176a9e6e62a',
    description: 'Moto',
    date: DateTime(2020, 1, 2),
    amountUsd: 542.96,
  );

  setUp(() async {
    await setUpTestHive();
    if (!Hive.isAdapterRegistered(0)) {
      Hive.registerAdapter(LocalTransactionAdapter());
    }
    service = LocalTransactionService();
  });

  tearDown(() async {
    await tearDownTestHive();
  });

  test('saveLocalTransaction deve salvar com o ID correto', () async {
    final fakeTransaction = LocalTransaction(
      id: 'a664d78d-cce6-4770-b287-b176a9e6e62a',
      description: 'Moto',
      date: DateTime(2020, 1, 2),
      amountUsd: 542.96,
    );

    await service.saveLocalTransaction(fakeTransaction);
    final box = await Hive.openBox<LocalTransaction>('local_transactions');
    final saved = box.get(fakeTransaction.id);
    expect(saved?.description, equals('Moto'));
  });

  test('deleteLocalTransaction deve remover pelo ID', () async {
    final box = await Hive.openBox<LocalTransaction>('local_transactions');
    await box.put(fakeTransaction.id, fakeTransaction);

    await service.deleteLocalTransaction(fakeTransaction.id);
    expect(box.get(fakeTransaction.id), isNull);
  });

  test('getAllLocalTransactions deve retornar lista de transações', () async {
    final box = await Hive.openBox<LocalTransaction>('local_transactions');
    await box.put(fakeTransaction.id, fakeTransaction);

    final result = await service.getAllLocalTransactions();
    expect(result, hasLength(1));
    expect(result.first.description, equals('Moto'));
  });
}
