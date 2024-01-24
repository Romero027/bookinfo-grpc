use proxy_wasm::traits::RootContext;
use proxy_wasm::traits::{Context, HttpContext};
use proxy_wasm::types::{Action, LogLevel};


#[no_mangle]
pub fn _start() {
    proxy_wasm::set_log_level(LogLevel::Trace);
    proxy_wasm::set_root_context(|_| -> Box<dyn RootContext> {
        Box::new(loggingRoot)
    });
    proxy_wasm::set_http_context(|context_id, _| -> Box<dyn HttpContext> {
        Box::new(loggingBody {
            context_id,
        })
    });
}

struct loggingRoot;

impl loggingBody {}

impl Context for loggingRoot {}

impl RootContext for loggingRoot {
    fn on_vm_start(&mut self, _: usize) -> bool {
        log::warn!("executing on_vm_start");
        true
    }
}

struct loggingBody {
    #[allow(unused)]
    context_id: u32,
}

impl Context for loggingBody {
    fn on_http_call_response(&mut self, _: u32, _: usize, body_size: usize, _: usize) {
        log::warn!(
            "executing on_http_call_response, self.context_id: {}",
            self.context_id
        );
    }
}

impl HttpContext for loggingBody {
    fn on_http_request_headers(&mut self, _num_of_headers: usize, _end_of_stream: bool) -> Action {
        log::warn!(
            "executing on_http_request_headers, self.context_id: {}",
            self.context_id
        );
        // if !end_of_stream {
        //     return Action::Continue;
        // }

        Action::Continue
    }

    fn on_http_request_body(&mut self, body_size: usize, _end_of_stream: bool) -> Action {
        log::warn!(
            "executing on_http_request_body, self.context_id: {}",
            self.context_id
        );


        Action::Continue
    }

    fn on_http_response_headers(&mut self, _num_headers: usize, _end_of_stream: bool) -> Action {
        log::warn!(
            "executing on_http_response_headers, self.context_id: {}",
            self.context_id
        );
        // if !end_of_stream {
        //    return Action::Continue;
        // }

        Action::Continue
    }

    fn on_http_response_body(&mut self, body_size: usize, _end_of_stream: bool) -> Action {
        log::warn!(
            "executing on_http_response_body, self.context_id: {}",
            self.context_id
        );
        // if !end_of_stream {
        //    return Action::Pause;
        // }


        Action::Continue
    }
}
