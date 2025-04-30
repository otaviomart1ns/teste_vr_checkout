import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/modules/transaction/create/store/transaction_create_store.dart';
import 'package:frontend/modules/transaction/create/transaction_create_page.dart';
import 'package:frontend/modules/transaction/pending/store/transaction_pending_store.dart';
import 'package:frontend/modules/transaction/pending/transaction_pending_page.dart';
import 'package:frontend/modules/transaction/shared/local_transaction_service.dart';
import 'package:frontend/modules/transaction/shared/transaction_service.dart';
import 'package:frontend/modules/transaction/view/store/transaction_view_store.dart';
import 'package:frontend/modules/transaction/view/transaction_view_page.dart';

class TransactionModule extends Module {
  @override
  void binds(Injector i) {
    // Shared services
    i.addLazySingleton(TransactionService.new);
    i.addLazySingleton(LocalTransactionService.new);

    // Stores
    i.addLazySingleton(
      () => TransactionCreateStore(
        i<TransactionService>(),
        i<LocalTransactionService>(),
      ),
    );
    i.addLazySingleton(() => TransactionViewStore(i<TransactionService>()));
    i.addLazySingleton(
      () => TransactionPendingStore(
        i<LocalTransactionService>(),
        i<TransactionService>(),
      ),
    );
  }

  @override
  void routes(RouteManager r) {
    r.child('/create', child: (context) => const TransactionCreatePage());
    r.child('/view', child: (context) => const TransactionViewPage());
    r.child('/pending', child: (context) => const TransactionPendingPage());
  }
}
