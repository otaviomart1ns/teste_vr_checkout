import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/modules/transaction/pending/transaction_pending_page.dart';
import 'package:frontend/modules/transaction/pending/store/transaction_pending_store.dart';
import 'package:frontend/modules/transaction/shared/local_transaction_service.dart';
import 'package:frontend/modules/transaction/shared/transaction_service.dart';

class TransactionPendingModule extends Module {
  @override
  void binds(Injector i) {
    i.addLazySingleton(TransactionService.new);
    i.addLazySingleton(LocalTransactionService.new);
    i.addLazySingleton(
      () => TransactionPendingStore(
        i.get<LocalTransactionService>(),
        i.get<TransactionService>(),
      ),
    );
  }

  @override
  void routes(RouteManager r) {
    r.child('/', child: (context) => const TransactionPendingPage());
  }
}
