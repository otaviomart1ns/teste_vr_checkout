import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/modules/transaction/view/store/transaction_view_store.dart';
import 'package:frontend/modules/transaction/view/transaction_view_page.dart';
import 'package:frontend/modules/transaction/shared/transaction_service.dart';

class TransactionViewModule extends Module {
  @override
  void binds(i) {
    i.addLazySingleton(TransactionService.new);
    i.addLazySingleton(() => TransactionViewStore(i<TransactionService>()));
  }

  @override
  void routes(r) {
    r.child('/', child: (context) => const TransactionViewPage());
  }
}
