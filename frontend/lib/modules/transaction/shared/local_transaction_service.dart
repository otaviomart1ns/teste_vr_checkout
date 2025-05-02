import 'package:frontend/modules/transaction/shared/models/local_transaction.dart';
import 'package:hive/hive.dart';

class LocalTransactionService {
  static const _boxName = 'local_transactions';

  Future<Box<LocalTransaction>> _openBox() async {
    return await Hive.openBox<LocalTransaction>(_boxName);
  }

  /// Salva uma transação localmente
  Future<void> saveLocalTransaction(LocalTransaction transaction) async {
    final box = await _openBox();
    await box.put(transaction.id, transaction);
  }

  /// Deleta uma transação local pelo ID
  Future<void> deleteLocalTransaction(String id) async {
    final box = await _openBox();
    await box.delete(id);
  }

  /// Lista transações
  Future<List<LocalTransaction>> getAllLocalTransactions() async {
    final box = await _openBox();
    return box.values.toList();
  }
}
