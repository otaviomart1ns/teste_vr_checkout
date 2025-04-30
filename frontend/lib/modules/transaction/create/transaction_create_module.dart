import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/modules/transaction/create/store/transaction_create_store.dart';
import 'package:frontend/modules/transaction/create/transaction_create_page.dart';
import 'package:frontend/modules/transaction/shared/local_transaction_service.dart';
import 'package:frontend/modules/transaction/shared/transaction_service.dart';

class TransactionCreateModule extends Module {
  @override
  void binds(i) {
    i.addLazySingleton(TransactionService.new);
    i.addLazySingleton(LocalTransactionService.new);
    i.addLazySingleton(
      () => TransactionCreateStore(
        i<TransactionService>(),
        i<LocalTransactionService>(),
      ),
    );
  }

  @override
  void routes(r) {
    r.child('/', child: (context) => const TransactionCreatePage());
  }
}
